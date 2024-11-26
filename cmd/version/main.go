/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package main

import (
	"fmt"

	"github.com/origadmin/toolkits/version"
)

func main() {
	// Call the Get function to retrieve the build information
	buildInfo := version.ReadBuildInfo()
	fmt.Println("Build Information:")
	fmt.Printf("Git Tag: %s\n", buildInfo.GitTag)
	fmt.Printf("Git Commit: %s\n", buildInfo.GitCommit)
	fmt.Printf("Git Branch: %s\n", buildInfo.GitBranch)
	fmt.Printf("Git Tree State: %s\n", buildInfo.GitTreeState)
	fmt.Printf("Build Date: %s\n", buildInfo.BuildDate)
	fmt.Printf("Built By: %s\n", buildInfo.BuiltBy)
	fmt.Printf("Go Version: %s\n", buildInfo.GoVersion)
	fmt.Printf("Compiler: %s\n", buildInfo.Compiler)
	fmt.Printf("Platform: %s\n", buildInfo.Platform)
	fmt.Printf("Version: %s\n", buildInfo.Version)

	// Call the Show function to print the build information in JSON format
	fmt.Println("\nBuild Information in JSON format:")
	version.PrintBuildInfo()
}
