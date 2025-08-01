// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package cmd defines a CLI command to init the project template.
package cmd

import (
	"github.com/spf13/cobra"
)

func InitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "init the project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
