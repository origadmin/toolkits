// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package config is the config package for origen
package config_test

import (
	"testing"

	"github.com/origadmin/toolkits/codec"

	"github.com/origadmin/origen/config"
)

func TestC(t *testing.T) {
	config := config.Config{
		Organization: "OrigAdmin",
		Application:  "Origen",
		Project:      "Origen",
		Version:      "v0.0.1",
		Type:         config.TypeProject,
		Mods: []config.Mod{
			{
				Name:   "",
				Repo:   "",
				Tag:    "",
				Branch: "",
				Commit: "",
			},
		},
		Platform: []config.Platform{
			{
				Name:    "backend",
				Version: "v0.0.1",
				Layers:  nil,
			},
		},
		Web: config.Web{
			TODO: "todo",
		},
		Static: "./static",
		Plugins: []config.Plugin{
			{
				Name: "name",
				TODO: "todo",
			},
		},
		Resources: config.Resources{
			Path: "./resources",
		},
		Document: config.Document{
			Title: "title",
		},
	}

	err := codec.EncodeTOMLFile("config.toml", config)
	if err != nil {
		t.Fatal(err)
	}
}
