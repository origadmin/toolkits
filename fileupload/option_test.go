/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package fileupload implements the functions, types, and interfaces for the module.
package fileupload

import (
	"testing"
)

func TestGenerateFileNameHash(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				name: "test.jpg",
			},
			want: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855.jpg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateFileNameHash(tt.args.name); got != tt.want {
				t.Errorf("GenerateFileNameHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
