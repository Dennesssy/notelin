package main

import (
	"fmt"
	"os"

	"github.com/notelin/notelin/pkg/ui"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "notelin",
		Short: "No-Tel-in: Telemetry & Sentry Management Scanner",
		Long: `No-Tel-in (Telemetry and Sentry Management)
Ain't no telling what data they be lyin about collectin.

A privacy-focused tool to discover, analyze, and manage telemetry 
and data collection from applications on your system.`,
		Version: fmt.Sprintf("%s (%s) built on %s", version, commit, date),
		RunE:    runInteractive,
	}

	var scanCmd = &cobra.Command{
		Use:   "scan",
		Short: "Scan system for applications with telemetry",
		RunE:  runScan,
	}

	var reportCmd = &cobra.Command{
		Use:   "report",
		Short: "Generate privacy assessment report",
		RunE:  runReport,
	}

	rootCmd.AddCommand(scanCmd, reportCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runInteractive(cmd *cobra.Command, args []string) error {
	app := ui.NewApp()
	return app.Run()
}

func runScan(cmd *cobra.Command, args []string) error {
	fmt.Println("🔍 Scanning system for suspicious applications...")
	// TODO: Implement scanning logic
	return nil
}

func runReport(cmd *cobra.Command, args []string) error {
	fmt.Println("📊 Generating privacy assessment report...")
	// TODO: Implement reporting logic
	return nil
}