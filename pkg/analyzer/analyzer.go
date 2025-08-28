package analyzer

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/notelin/notelin/pkg/scanner"
)

// Analyzer provides detailed analysis of applications
type Analyzer struct {
	homeDir string
}

// PrivacyReport contains detailed privacy assessment
type PrivacyReport struct {
	Application     scanner.Application `json:"application"`
	PrivacyScore    int                 `json:"privacy_score"`
	DataCategories  []DataCategory      `json:"data_categories"`
	Recommendations []Recommendation    `json:"recommendations"`
	CleanupActions  []CleanupAction     `json:"cleanup_actions"`
}

// DataCategory represents a type of data being collected
type DataCategory struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Severity    string   `json:"severity"`
	Evidence    []string `json:"evidence"`
	Impact      string   `json:"impact"`
}

// Recommendation suggests privacy improvements
type Recommendation struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Urgency     string `json:"urgency"`
	Action      string `json:"action"`
}

// CleanupAction represents actionable cleanup steps
type CleanupAction struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Paths       []string `json:"paths"`
	Safe        bool     `json:"safe"`
	SpaceFreed  int64    `json:"space_freed"`
}

// NewAnalyzer creates a new analyzer instance
func NewAnalyzer() *Analyzer {
	return &Analyzer{
		homeDir: os.Getenv("HOME"),
	}
}

// AnalyzeApplication performs deep privacy analysis
func (a *Analyzer) AnalyzeApplication(app scanner.Application) (*PrivacyReport, error) {
	report := &PrivacyReport{
		Application:     app,
		DataCategories:  []DataCategory{},
		Recommendations: []Recommendation{},
		CleanupActions:  []CleanupAction{},
	}

	// Analyze data collection patterns
	a.analyzeDataCategories(report)
	
	// Calculate privacy score
	a.calculatePrivacyScore(report)
	
	// Generate recommendations
	a.generateRecommendations(report)
	
	// Find cleanup opportunities
	a.identifyCleanupActions(report)

	return report, nil
}

// analyzeDataCategories identifies specific data collection categories
func (a *Analyzer) analyzeDataCategories(report *PrivacyReport) {
	app := report.Application
	supportPath := filepath.Join(a.homeDir, "Library", "Application Support", app.Name)

	// Session/Usage Tracking
	if a.hasSessionTracking(app, supportPath) {
		report.DataCategories = append(report.DataCategories, DataCategory{
			Name:        "Session Tracking",
			Description: "Tracks user sessions, window states, and usage patterns",
			Severity:    "Medium",
			Evidence:    []string{"window-state.json", "recent-dirs.json", "session data"},
			Impact:      "Reveals workflow patterns and frequently accessed locations",
		})
	}

	// Environment Data
	if a.hasEnvironmentLogging(app, supportPath) {
		report.DataCategories = append(report.DataCategories, DataCategory{
			Name:        "Environment Logging",
			Description: "Records system environment variables and paths",
			Severity:    "High", 
			Evidence:    []string{"logs/main.log", "environment variables"},
			Impact:      "Exposes system configuration and potentially sensitive paths",
		})
	}

	// Network Persistence
	if a.hasNetworkPersistence(app, supportPath) {
		report.DataCategories = append(report.DataCategories, DataCategory{
			Name:        "Network State Persistence",
			Description: "Stores network connection data and trust tokens",
			Severity:    "Medium",
			Evidence:    []string{"Network Persistent State", "Trust Tokens"},
			Impact:      "Could reveal connection patterns and remote services used",
		})
	}

	// Browser Engine Data (for Electron apps)
	if a.isElectronApp(app) {
		report.DataCategories = append(report.DataCategories, DataCategory{
			Name:        "Browser Engine Data",
			Description: "Chromium-based cache, cookies, and local storage",
			Severity:    "Medium",
			Evidence:    []string{"Cache/", "Local Storage/", "Cookies"},
			Impact:      "Standard browser tracking mechanisms available to app",
		})
	}

	// Remote Configuration
	if a.hasRemoteConfig(app, supportPath) {
		report.DataCategories = append(report.DataCategories, DataCategory{
			Name:        "Remote Configuration",
			Description: "App behavior controlled by external servers",
			Severity:    "High",
			Evidence:    []string{"remote_config_data", "Firebase config"},
			Impact:      "App features and behavior can be modified without updates",
		})
	}
}

// calculatePrivacyScore computes overall privacy risk (0-100, higher = worse)
func (a *Analyzer) calculatePrivacyScore(report *PrivacyReport) {
	score := 0
	
	for _, category := range report.DataCategories {
		switch category.Severity {
		case "High":
			score += 25
		case "Medium":
			score += 15
		case "Low":
			score += 5
		}
	}

	// Additional scoring based on app characteristics
	app := report.Application
	if app.Source != "App Store" {
		score += 10 // Non-App Store apps have fewer restrictions
	}

	if len(app.DataSources) > 3 {
		score += 15 // Multiple data collection mechanisms
	}

	// Cap at 100
	if score > 100 {
		score = 100
	}

	report.PrivacyScore = score
}

// generateRecommendations creates actionable privacy recommendations
func (a *Analyzer) generateRecommendations(report *PrivacyReport) {
	app := report.Application

	if report.PrivacyScore > 70 {
		report.Recommendations = append(report.Recommendations, Recommendation{
			Type:        "critical",
			Description: "Consider removing this application due to extensive data collection",
			Urgency:     "High",
			Action:      "Evaluate if app functionality justifies privacy trade-offs",
		})
	}

	// Specific recommendations based on data categories
	for _, category := range report.DataCategories {
		switch category.Name {
		case "Remote Configuration":
			report.Recommendations = append(report.Recommendations, Recommendation{
				Type:        "security",
				Description: "Block network access to prevent remote configuration updates",
				Urgency:     "Medium", 
				Action:      "Use firewall rules or Little Snitch to block external connections",
			})
		case "Environment Logging":
			report.Recommendations = append(report.Recommendations, Recommendation{
				Type:        "privacy",
				Description: "Clear environment logs that may contain sensitive paths",
				Urgency:     "Medium",
				Action:      "Regularly clean log files in Application Support directory",
			})
		case "Session Tracking":
			report.Recommendations = append(report.Recommendations, Recommendation{
				Type:        "privacy",
				Description: "Clear session data to remove usage pattern history",
				Urgency:     "Low",
				Action:      "Delete recent directories and window state files",
			})
		}
	}

	// App-specific recommendations
	appNameLower := strings.ToLower(app.Name)
	if strings.Contains(appNameLower, "nicegram") {
		report.Recommendations = append(report.Recommendations, Recommendation{
			Type:        "critical",
			Description: "Nicegram Desktop collects extensive analytics despite privacy claims",
			Urgency:     "High", 
			Action:      "Consider using official Telegram client or Signal instead",
		})
	}
}

// identifyCleanupActions finds data that can be safely cleaned
func (a *Analyzer) identifyCleanupActions(report *PrivacyReport) {
	app := report.Application
	supportPath := filepath.Join(a.homeDir, "Library", "Application Support", app.Name)

	// Cache cleanup
	cachePaths := []string{
		filepath.Join(supportPath, "Cache"),
		filepath.Join(supportPath, "Partitions", app.Name, "Cache"),
		filepath.Join(supportPath, "Code Cache"),
		filepath.Join(supportPath, "GPUCache"),
	}

	var existingCaches []string
	var totalCacheSize int64

	for _, path := range cachePaths {
		if size := a.getDirSize(path); size > 0 {
			existingCaches = append(existingCaches, path)
			totalCacheSize += size
		}
	}

	if len(existingCaches) > 0 {
		report.CleanupActions = append(report.CleanupActions, CleanupAction{
			Name:        "Clear Application Cache",
			Description: "Remove cached data to free space and reduce tracking",
			Paths:       existingCaches,
			Safe:        true,
			SpaceFreed:  totalCacheSize,
		})
	}

	// Session data cleanup
	sessionPaths := []string{
		filepath.Join(supportPath, "recent-dirs.json"),
		filepath.Join(supportPath, "window-state.json"),
		filepath.Join(supportPath, "Local Storage"),
		filepath.Join(supportPath, "Session Storage"),
	}

	var existingSessions []string
	for _, path := range sessionPaths {
		if a.pathExists(path) {
			existingSessions = append(existingSessions, path)
		}
	}

	if len(existingSessions) > 0 {
		report.CleanupActions = append(report.CleanupActions, CleanupAction{
			Name:        "Clear Session Data",
			Description: "Remove usage history and session persistence",
			Paths:       existingSessions,
			Safe:        true,
			SpaceFreed:  a.getPathsSize(existingSessions),
		})
	}

	// Log cleanup (be careful - may affect app functionality)
	logPath := filepath.Join(supportPath, "logs")
	if a.pathExists(logPath) {
		report.CleanupActions = append(report.CleanupActions, CleanupAction{
			Name:        "Clear Application Logs",
			Description: "Remove logs that may contain sensitive information",
			Paths:       []string{logPath},
			Safe:        false, // May affect debugging/support
			SpaceFreed:  a.getDirSize(logPath),
		})
	}
}

// Helper methods

func (a *Analyzer) hasSessionTracking(app scanner.Application, supportPath string) bool {
	patterns := []string{"recent-dirs.json", "window-state.json", "sessions"}
	return a.hasAnyPattern(supportPath, patterns)
}

func (a *Analyzer) hasEnvironmentLogging(app scanner.Application, supportPath string) bool {
	logPath := filepath.Join(supportPath, "logs", "main.log")
	return a.pathExists(logPath)
}

func (a *Analyzer) hasNetworkPersistence(app scanner.Application, supportPath string) bool {
	patterns := []string{"Network Persistent State", "Trust Tokens"}
	return a.hasAnyPattern(supportPath, patterns)
}

func (a *Analyzer) hasRemoteConfig(app scanner.Application, supportPath string) bool {
	patterns := []string{"remote_config_data", "__FIRAPP_DEFAULT", "firebase"}
	return a.hasAnyPattern(supportPath, patterns)
}

func (a *Analyzer) isElectronApp(app scanner.Application) bool {
	return app.Metadata["framework"] == "electron" ||
		a.pathExists(filepath.Join(app.Path, "Contents", "Frameworks", "Electron Framework.framework"))
}

func (a *Analyzer) hasAnyPattern(basePath string, patterns []string) bool {
	for _, pattern := range patterns {
		if a.pathExists(filepath.Join(basePath, pattern)) {
			return true
		}
	}
	return false
}

func (a *Analyzer) pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (a *Analyzer) getDirSize(path string) int64 {
	var size int64
	filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size
}

func (a *Analyzer) getPathsSize(paths []string) int64 {
	var totalSize int64
	for _, path := range paths {
		if info, err := os.Stat(path); err == nil {
			if info.IsDir() {
				totalSize += a.getDirSize(path)
			} else {
				totalSize += info.Size()
			}
		}
	}
	return totalSize
}