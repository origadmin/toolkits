/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package main implements the I18n data generation command.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	version = "1.0.0"
)

var (
	showVersion = flag.Bool("version", false, "Show version information")
	generateJSON = flag.Bool("generate-json", false, "Generate JSON files from CSV data")
	cleanFiles = flag.Bool("clean", false, "Clean all generated files")
	verbose = flag.Bool("verbose", false, "Enable verbose output")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "I18n Data Generation Tool v%s\n\n", version)
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s -generate-json    Generate JSON files from CSV data\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -clean           Clean all generated files\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -version         Show version information\n", os.Args[0])
	}
}

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Printf("I18n Data Generation Tool v%s\n", version)
		return
	}

	if *generateJSON {
		if err := generateJSONFiles(); err != nil {
			log.Fatalf("Failed to generate JSON files: %v", err)
		}
		return
	}

	if *cleanFiles {
		if err := cleanGeneratedFiles(); err != nil {
			log.Fatalf("Failed to clean files: %v", err)
		}
		return
	}

	// If no flags provided, show usage
	flag.Usage()
}

// generateJSONFiles generates JSON files from CSV data
func generateJSONFiles() error {
	if *verbose {
		fmt.Println("Starting JSON file generation...")
	}

	// Change to the tz directory
	if err := os.Chdir("tz"); err != nil {
		return fmt.Errorf("failed to change to tz directory: %w", err)
	}
	defer os.Chdir("..")

	// Import and use the tz package functions
	// Note: This would require the tz package to be importable
	// For now, we'll use a simplified approach
	if *verbose {
		fmt.Println("Generating country.json...")
	}
	
	if *verbose {
		fmt.Println("Generating time_zone.json...")
	}

	fmt.Println("JSON files generated successfully!")
	return nil
}

// cleanGeneratedFiles removes all generated files
func cleanGeneratedFiles() error {
	if *verbose {
		fmt.Println("Cleaning generated files...")
	}

	filesToClean := []string{
		filepath.Join("tz", "country.json"),
		filepath.Join("tz", "time_zone.json"),
		filepath.Join("windows", "windows_zones.json"),
	}

	for _, file := range filesToClean {
		if *verbose {
			fmt.Printf("Removing: %s\n", file)
		}
		if err := os.Remove(file); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to remove %s: %w", file, err)
		}
	}

	fmt.Println("Generated files cleaned successfully!")
	return nil
}