// Copyright (c) 2024 OrigAdmin. All rights reserved.

//go:generate swag init --parseDependency --generalInfo ./main.go --output ./docs
// #go:generate docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate --git-repo-id swagger-api --git-user-id origadmin -i /local/docs/swagger.json -g go-gin-server -o /local/docs/v3

// Package main is the main package for origen
package main

import (
	"fmt"
	"os"

	goversion "github.com/caarlos0/go-version"
	"github.com/spf13/cobra"

	"github.com/origadmin/origen/cmd"
	"github.com/origadmin/origen/config"
)

// tags for goreleaser
var (
	version   = ""
	commit    = ""
	treeState = ""
	date      = ""
	builtBy   = ""
	// debug     = false
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "origen",
	Short: "A scaffolding for quickly setting up a project.",
	Long: "A scaffolding for quickly setting up a project.\n" +
		"Includes but is not limited to RBAC, user management,\n" +
		"logging, file management, and more.",
	Run: func(cmd *cobra.Command, args []string) {
		// Place your logic here
		// _ = cmd.Help()
	},
}

func init() {
	// cobra.AddTemplateFuncs(usage.Template)
	goinfo := buildVersion(version, commit, date, builtBy, treeState)
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.CompletionOptions.DisableNoDescFlag = true
	// rootCmd.SetUsageTemplate(resources.Usage)
	rootCmd.AddCommand(cmd.InitCmd(), cmd.NewCmd())
	rootCmd.Version = goinfo.String()
}

func main() {
	Execute()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Command executed with error:\n%v\n", err)
		os.Exit(1)
	}
}

func buildVersion(version, commit, date, builtBy, treeState string) goversion.Info {
	return goversion.GetVersionInfo(
		goversion.WithAppDetails(config.Project, "A scaffolding for quickly setting up a project.", config.WebSite),
		goversion.WithASCIIName(config.UI),
		func(i *goversion.Info) {
			if commit != "" {
				i.GitCommit = commit
			}
			if version != "" {
				i.GitVersion = version
			}
			if treeState != "" {
				i.GitTreeState = treeState
			}
			if date != "" {
				i.BuildDate = date
			}
			if builtBy != "" {
				i.BuiltBy = builtBy
			}
		},
	)
}
