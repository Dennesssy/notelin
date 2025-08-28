package scanner

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Application represents a discovered application
type Application struct {
	Name         string            `json:"name"`
	Path         string            `json:"path"`
	Version      string            `json:"version"`
	Developer    string            `json:"developer"`
	Source       string            `json:"source"` // App Store, Identified Developer, etc.
	RiskLevel    RiskLevel         `json:"risk_level"`
	DataSources  []DataSource      `json:"data_sources"`
	Permissions  []string          `json:"permissions"`
	NetworkHosts []string          `json:"network_hosts"`
	LastModified time.Time         `json:"last_modified"`
	Metadata     map[string]string `json:"metadata"`
}

// RiskLevel represents the privacy risk assessment
type RiskLevel string

const (
	RiskLow    RiskLevel = "low"
	RiskMedium RiskLevel = "medium" 
	RiskHigh   RiskLevel = "high"
)

// DataSource represents detected data collection
type DataSource struct {
	Type        string   `json:"type"`        // firebase, sentry, analytics, etc.
	Description string   `json:"description"`
	Evidence    []string `json:"evidence"`    // file paths, config entries, etc.
	OptOut      bool     `json:"opt_out"`     // whether opt-out is available
}

// Scanner handles application discovery and analysis
type Scanner struct {
	appPaths   []string
	knownRisks map[string]RiskLevel
}

// NewScanner creates a new scanner instance
func NewScanner() *Scanner {
	return &Scanner{
		appPaths: []string{
			"/Applications",
			"/System/Applications", 
			filepath.Join(os.Getenv("HOME"), "Applications"),
		},
		knownRisks: map[string]RiskLevel{
			"nicegram":     RiskHigh,
			"spotify":      RiskHigh,
			"zoom":         RiskMedium,
			"blackbox":     RiskMedium,
			"adobe":        RiskMedium,
			"vscode":       RiskLow,
			"unarchiver":   RiskLow,
		},
	}
}

// ScanApplications discovers and analyzes applications
func (s *Scanner) ScanApplications() ([]Application, error) {
	var applications []Application

	for _, appPath := range s.appPaths {
		apps, err := s.scanPath(appPath)
		if err != nil {
			continue // Skip inaccessible paths
		}
		applications = append(applications, apps...)
	}

	// Filter out Apple applications
	var nonAppleApps []Application
	for _, app := range applications {
		if !s.isAppleApp(app) {
			nonAppleApps = append(nonAppleApps, app)
		}
	}

	return nonAppleApps, nil
}

// scanPath scans a specific directory for applications
func (s *Scanner) scanPath(path string) ([]Application, error) {
	var apps []Application

	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil // Skip errors, continue scanning
		}

		if d.IsDir() && strings.HasSuffix(path, ".app") {
			app, err := s.analyzeApplication(path)
			if err == nil {
				apps = append(apps, *app)
			}
			return filepath.SkipDir // Don't scan inside .app bundles
		}

		return nil
	})

	return apps, err
}

// analyzeApplication performs detailed analysis of an application
func (s *Scanner) analyzeApplication(appPath string) (*Application, error) {
	app := &Application{
		Path:        appPath,
		Name:        s.extractAppName(appPath),
		DataSources: []DataSource{},
		Permissions: []string{},
		Metadata:    make(map[string]string),
	}

	// Get basic app info
	if err := s.getAppInfo(app); err != nil {
		return nil, err
	}

	// Analyze data collection patterns
	s.analyzeDataCollection(app)

	// Assess risk level
	s.assessRisk(app)

	return app, nil
}

// getAppInfo extracts basic application information
func (s *Scanner) getAppInfo(app *Application) error {
	// Try to read Info.plist
	infoPlistPath := filepath.Join(app.Path, "Contents", "Info.plist")
	
	// For now, we'll use basic file info
	if info, err := os.Stat(app.Path); err == nil {
		app.LastModified = info.ModTime()
	}

	// Extract version and developer from system_profiler if available
	// This would normally parse plist files
	app.Version = "unknown"
	app.Developer = "unknown"
	app.Source = "Unknown"

	return nil
}

// analyzeDataCollection looks for data collection evidence
func (s *Scanner) analyzeDataCollection(app *Application) {
	appSupportPath := filepath.Join(os.Getenv("HOME"), "Library", "Application Support", app.Name)
	
	// Check for common data collection patterns
	s.checkFirebase(app, appSupportPath)
	s.checkSentry(app, appSupportPath)
	s.checkAnalytics(app, appSupportPath)
	s.checkElectron(app)
}

// checkFirebase looks for Firebase-related data collection
func (s *Scanner) checkFirebase(app *Application, supportPath string) {
	patterns := []string{
		"remote_config_data",
		"firebase",
		"__FIRAPP_DEFAULT",
	}

	for _, pattern := range patterns {
		if s.pathExists(filepath.Join(supportPath, pattern)) {
			app.DataSources = append(app.DataSources, DataSource{
				Type:        "firebase",
				Description: "Firebase Remote Config - Can modify app behavior remotely",
				Evidence:    []string{filepath.Join(supportPath, pattern)},
				OptOut:      false,
			})
			break
		}
	}
}

// checkSentry looks for crash reporting and monitoring
func (s *Scanner) checkSentry(app *Application, supportPath string) {
	// Check for Sentry in logs or config
	if strings.Contains(strings.ToLower(app.Name), "sentry") ||
		s.pathExists(filepath.Join(supportPath, "Crashpad")) {
		app.DataSources = append(app.DataSources, DataSource{
			Type:        "sentry", 
			Description: "Crash reporting and error monitoring",
			Evidence:    []string{filepath.Join(supportPath, "Crashpad")},
			OptOut:      false,
		})
	}
}

// checkAnalytics looks for analytics and tracking
func (s *Scanner) checkAnalytics(app *Application, supportPath string) {
	analyticsPatterns := []string{
		"analytics",
		"telemetry",
		"usage",
		"metrics",
	}

	for _, pattern := range analyticsPatterns {
		searchPath := filepath.Join(supportPath, pattern)
		if s.pathExists(searchPath) {
			app.DataSources = append(app.DataSources, DataSource{
				Type:        "analytics",
				Description: "Usage analytics and telemetry data collection", 
				Evidence:    []string{searchPath},
				OptOut:      true, // Many apps allow opting out of analytics
			})
			break
		}
	}
}

// checkElectron detects Electron-based apps which typically have more telemetry
func (s *Scanner) checkElectron(app *Application) {
	electronPaths := []string{
		filepath.Join(app.Path, "Contents", "Frameworks", "Electron Framework.framework"),
		filepath.Join(app.Path, "Contents", "Resources", "app.asar"),
	}

	for _, path := range electronPaths {
		if s.pathExists(path) {
			app.DataSources = append(app.DataSources, DataSource{
				Type:        "electron",
				Description: "Electron framework - Chromium-based data collection",
				Evidence:    []string{path},
				OptOut:      false,
			})
			app.Metadata["framework"] = "electron"
			break
		}
	}
}

// assessRisk determines the overall privacy risk level
func (s *Scanner) assessRisk(app *Application) {
	appNameLower := strings.ToLower(app.Name)
	
	// Check known risks first
	for name, risk := range s.knownRisks {
		if strings.Contains(appNameLower, name) {
			app.RiskLevel = risk
			return
		}
	}

	// Calculate risk based on data sources
	riskScore := 0
	for _, ds := range app.DataSources {
		switch ds.Type {
		case "firebase":
			riskScore += 3
		case "sentry":
			riskScore += 2
		case "analytics":
			if ds.OptOut {
				riskScore += 1
			} else {
				riskScore += 2
			}
		case "electron":
			riskScore += 1
		}
	}

	if riskScore >= 4 {
		app.RiskLevel = RiskHigh
	} else if riskScore >= 2 {
		app.RiskLevel = RiskMedium
	} else {
		app.RiskLevel = RiskLow
	}
}

// isAppleApp determines if an application is from Apple
func (s *Scanner) isAppleApp(app Application) bool {
	applePrefixes := []string{
		"com.apple.",
		"/System/Applications",
	}

	for _, prefix := range applePrefixes {
		if strings.HasPrefix(app.Path, prefix) ||
			strings.Contains(strings.ToLower(app.Developer), "apple") {
			return true
		}
	}

	return false
}

// extractAppName extracts the application name from its path
func (s *Scanner) extractAppName(appPath string) string {
	base := filepath.Base(appPath)
	return strings.TrimSuffix(base, ".app")
}

// pathExists checks if a path exists
func (s *Scanner) pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// ExportResults exports scan results to JSON
func (s *Scanner) ExportResults(apps []Application, outputPath string) error {
	data, err := json.MarshalIndent(apps, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(outputPath, data, 0644)
}