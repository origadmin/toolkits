package argon2

import (
	"testing"
)

func TestParams_ParseAndString(t *testing.T) {
	tests := []struct {
		name     string
		params   string
		wantErr  bool
		validate func(*testing.T, *Params)
	}{
		{
			name:    "完整参数",
			params:  "t:3,m:65536,p:4,k:32",
			wantErr: false,
			validate: func(t *testing.T, p *Params) {
				if p.TimeCost != 3 {
					t.Errorf("TimeCost = %v, want %v", p.TimeCost, 3)
				}
				if p.MemoryCost != 65536 {
					t.Errorf("MemoryCost = %v, want %v", p.MemoryCost, 65536)
				}
				if p.Threads != 4 {
					t.Errorf("Threads = %v, want %v", p.Threads, 4)
				}
				if p.KeyLength != 32 {
					t.Errorf("KeyLength = %v, want %v", p.KeyLength, 32)
				}
			},
		},
		{
			name:    "部分参数",
			params:  "t:3,m:65536",
			wantErr: false,
			validate: func(t *testing.T, p *Params) {
				if p.TimeCost != 3 {
					t.Errorf("TimeCost = %v, want %v", p.TimeCost, 3)
				}
				if p.MemoryCost != 65536 {
					t.Errorf("MemoryCost = %v, want %v", p.MemoryCost, 65536)
				}
				if p.Threads != 0 {
					t.Errorf("Threads = %v, want %v", p.Threads, 0)
				}
				if p.KeyLength != 0 {
					t.Errorf("KeyLength = %v, want %v", p.KeyLength, 0)
				}
			},
		},
		{
			name:    "空参数",
			params:  "",
			wantErr: false,
			validate: func(t *testing.T, p *Params) {
				if p.TimeCost != 0 {
					t.Errorf("TimeCost = %v, want %v", p.TimeCost, 0)
				}
				if p.MemoryCost != 0 {
					t.Errorf("MemoryCost = %v, want %v", p.MemoryCost, 0)
				}
				if p.Threads != 0 {
					t.Errorf("Threads = %v, want %v", p.Threads, 0)
				}
				if p.KeyLength != 0 {
					t.Errorf("KeyLength = %v, want %v", p.KeyLength, 0)
				}
			},
		},
		{
			name:    "无效参数格式",
			params:  "t:3,m:65536,p:4,k:32,invalid",
			wantErr: true,
		},
		{
			name:    "无效参数值",
			params:  "t:invalid,m:65536",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params, err := parseParams(tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if tt.validate != nil {
				tt.validate(t, params)
			}

			// 测试 String 方法
			str := params.String()
			if str != tt.params && tt.params != "" {
				t.Errorf("String() = %v, want %v", str, tt.params)
			}
		})
	}
}

func TestParams_String(t *testing.T) {
	tests := []struct {
		name   string
		params *Params
		want   string
	}{
		{
			name: "完整参数",
			params: &Params{
				TimeCost:   3,
				MemoryCost: 65536,
				Threads:    4,
				KeyLength:  32,
			},
			want: "t:3,m:65536,p:4,k:32",
		},
		{
			name: "部分参数",
			params: &Params{
				TimeCost:   3,
				MemoryCost: 65536,
			},
			want: "t:3,m:65536",
		},
		{
			name:   "零值参数",
			params: &Params{},
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.params.String(); got != tt.want {
				t.Errorf("Params.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
