// Copyright (c) 2024 OrigAdmin. All rights reserved.
package io

import (
	"testing"
)

func TestDeleteFile(t *testing.T) {
	type args struct {
		path  string
		param string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				path:  "D:\\workspace\\project\\sugoitech\\admin\\data\\upload\\01HVVE64TWS98M1ZFYAYFXWHHM",
				param: "01HWAXGRX7ABY8BN353K7N12XY",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteFile(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
