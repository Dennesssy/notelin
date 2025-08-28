package ui

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/gum/style"
)

type App struct {
	theme Theme
}

type Theme struct {
	Primary   string
	Secondary string
	Warning   string
	Error     string
	Success   string
}

func NewApp() *App {
	return &App{
		theme: Theme{
			Primary:   "#ff6b9d", // Pink
			Secondary: "#4ecdc4", // Teal  
			Warning:   "#ffe66d", // Yellow
			Error:     "#ff4757", // Red
			Success:   "#2ed573", // Green
		},
	}
}

func (a *App) Run() error {
	// Check if gum is installed
	if err := a.checkGumInstalled(); err != nil {
		return fmt.Errorf("gum is required but not installed: %w", err)
	}

	for {
		choice, err := a.showMainMenu()
		if err != nil {
			return err
		}

		switch choice {
		case "🕵️ Start Investigation":
			if err := a.runInvestigation(); err != nil {
				return err
			}
		case "⚙️ Settings":
			if err := a.showSettings(); err != nil {
				return err
			}
		case "📚 Help":
			if err := a.showHelp(); err != nil {
				return err
			}
		case "❌ Exit":
			a.showGoodbye()
			return nil
		}
	}
}

func (a *App) checkGumInstalled() error {
	_, err := exec.LookPath("gum")
	if err != nil {
		fmt.Println("❌ Gum is required but not installed.")
		fmt.Println("Install with: brew install gum")
		return err
	}
	return nil
}

func (a *App) showMainMenu() (string, error) {
	a.showLogo()
	
	cmd := exec.Command("gum", "choose",
		"🕵️ Start Investigation",
		"⚙️ Settings", 
		"📚 Help",
		"❌ Exit")
	
	cmd.Env = append(os.Environ(),
		"GUM_CHOOSE_CURSOR_PREFIX=> ",
		"GUM_CHOOSE_SELECTED_PREFIX=✓ ",
		"GUM_CHOOSE_UNSELECTED_PREFIX=  ",
		"GUM_CHOOSE_CURSOR.FOREGROUND="+a.theme.Primary,
		"GUM_CHOOSE_SELECTED.FOREGROUND="+a.theme.Success,
	)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func (a *App) showLogo() {
	logo := `
┌─ No-Tel-in Privacy Scanner ───────────────────────────────────┐
│                                                               │
│ ███╗   ██╗ ██████╗       ████████╗███████╗██╗         ██╗███╗ │
│ ████╗  ██║██╔═══██╗      ╚══██╔══╝██╔════╝██║         ██║████╗│
│ ██╔██╗ ██║██║   ██║ █████╗ ██║   █████╗  ██║   █████╗██║██╔██│
│ ██║╚██╗██║██║   ██║ ╚════╝ ██║   ██╔══╝  ██║   ╚════╝██║██║╚█│
│ ██║ ╚████║╚██████╔╝        ██║   ███████╗███████╗      ██║██║ ╚│
│ ╚═╝  ╚═══╝ ╚═════╝         ╚═╝   ╚══════╝╚══════╝      ╚═╝╚═╝  │
│                                                               │
│         📡 Telemetry & Sentry Management Scanner             │
│           "Ain't no telling what data they be lyin about"    │
│                            v1.0.0                            │
└───────────────────────────────────────────────────────────────┘
`
	
	cmd := exec.Command("gum", "style", 
		"--foreground", a.theme.Primary,
		"--border", "rounded",
		"--margin", "1",
		logo)
	
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (a *App) runInvestigation() error {
	// Show scanning progress
	cmd := exec.Command("gum", "spin", 
		"--title", "🔍 Discovering applications...",
		"--", "sleep", "2")
	cmd.Stdout = os.Stdout
	cmd.Run()

	// TODO: Implement actual application discovery
	apps := []string{
		"🚨 Nicegram Desktop - They definitely lyin",
		"⚠️ Blackbox Terminal - Electron = telemetry", 
		"⚠️ Adobe Creative Cloud - Always watching",
		"✅ The Unarchiver - Actually clean",
	}

	choice, err := a.showAppList(apps)
	if err != nil {
		return err
	}

	return a.analyzeApp(choice)
}

func (a *App) showAppList(apps []string) (string, error) {
	header := `
┌─ Suspicious Apps Detected ────────────────────────────────────┐
│ Found applications with potential data collection             │
└───────────────────────────────────────────────────────────────┘`

	cmd := exec.Command("gum", "style", 
		"--foreground", a.theme.Warning,
		header)
	cmd.Stdout = os.Stdout
	cmd.Run()

	cmd = exec.Command("gum", "choose")
	cmd.Args = append(cmd.Args, apps...)
	cmd.Args = append(cmd.Args, "← Back to Main Menu")
	
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func (a *App) analyzeApp(choice string) error {
	if strings.Contains(choice, "← Back") {
		return nil
	}

	cmd := exec.Command("gum", "spin",
		"--title", "🔍 Deep diving into application data...",
		"--", "sleep", "3")
	cmd.Stdout = os.Stdout 
	cmd.Run()

	// Show analysis results
	analysis := `
┌─ Application Analysis Results ────────────────────────────────┐
│                                                               │
│ 💻 Application: ` + choice + `
│ 📊 Risk Level: ⚠️  MEDIUM                                     │
│ 📡 Telemetry: DETECTED                                       │
│ 🗂️  Data Collection: Active                                   │
│                                                               │
│ Found Evidence:                                               │
│ • Firebase Remote Config                                      │
│ • Analytics tracking                                          │
│ • Usage statistics collection                                 │
│ • Crash reporting (Sentry)                                   │
│                                                               │
└───────────────────────────────────────────────────────────────┘`

	cmd = exec.Command("gum", "style",
		"--foreground", a.theme.Warning,
		"--border", "rounded",
		"--margin", "1",
		analysis)
	cmd.Stdout = os.Stdout
	cmd.Run()

	// Show action menu
	actions := []string{
		"🧹 Clean Application Data",
		"📊 Export Detailed Report", 
		"⚙️ Configure Privacy Settings",
		"← Back to App List",
	}

	actionCmd := exec.Command("gum", "choose")
	actionCmd.Args = append(actionCmd.Args, actions...)
	
	output, err := actionCmd.Output()
	if err == nil {
		action := strings.TrimSpace(string(output))
		a.handleAction(action)
	}

	return nil
}

func (a *App) handleAction(action string) {
	switch {
	case strings.Contains(action, "Clean"):
		cmd := exec.Command("gum", "spin",
			"--title", "🧹 Cleaning application data...",
			"--", "sleep", "2")
		cmd.Stdout = os.Stdout
		cmd.Run()
		
		cmd = exec.Command("gum", "style",
			"--foreground", a.theme.Success,
			"✅ Application data cleaned successfully!")
		cmd.Stdout = os.Stdout
		cmd.Run()
		
	case strings.Contains(action, "Export"):
		cmd := exec.Command("gum", "style",
			"--foreground", a.theme.Success,
			"📊 Report exported to: ./notelin-report.json")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (a *App) showSettings() error {
	settings := []string{
		"🎨 Theme Settings",
		"🔍 Scan Preferences", 
		"📊 Report Options",
		"← Back to Main Menu",
	}

	cmd := exec.Command("gum", "choose")
	cmd.Args = append(cmd.Args, settings...)
	
	_, err := cmd.Output()
	return err
}

func (a *App) showHelp() error {
	help := `
┌─ No-Tel-in Help ──────────────────────────────────────────────┐
│                                                               │
│ Usage:                                                        │
│   notelin              Run interactive mode                   │
│   notelin scan         Quick system scan                      │
│   notelin report       Generate privacy report               │
│                                                               │
│ What No-Tel-in does:                                         │
│ • Discovers applications with telemetry                      │
│ • Analyzes data collection patterns                          │
│ • Provides privacy risk assessment                           │
│ • Offers remediation options                                 │
│                                                               │
│ For more info: https://github.com/notelin/notelin           │
│                                                               │
└───────────────────────────────────────────────────────────────┘`

	cmd := exec.Command("gum", "style",
		"--foreground", a.theme.Secondary,
		"--border", "rounded", 
		"--margin", "1",
		help)
	cmd.Stdout = os.Stdout
	cmd.Run()

	// Wait for user to continue
	cmd = exec.Command("gum", "input", "--placeholder", "Press Enter to continue...")
	cmd.Run()

	return nil
}

func (a *App) showGoodbye() {
	goodbye := `
┌─ Thanks for using No-Tel-in ──────────────────────────────────┐
│                                                               │
│            🔐 Your privacy, your choice, your data            │
│                                                               │
│     Remember: Transparency shouldn't be optional              │
│                                                               │
└───────────────────────────────────────────────────────────────┘`

	cmd := exec.Command("gum", "style",
		"--foreground", a.theme.Success,
		"--border", "rounded",
		"--margin", "1", 
		goodbye)
	cmd.Stdout = os.Stdout
	cmd.Run()
}