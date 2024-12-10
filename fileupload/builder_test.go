/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package fileupload implements the functions, types, and interfaces for the module.
package fileupload

import (
	"testing"
)

func TestGenerateRandomHash(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomHash(); len(got) != tt.want {
				t.Errorf("GenerateRandomHash() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
