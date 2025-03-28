package argon2

import (
	"testing"

	"github.com/origadmin/toolkits/crypto/hash/types"
)

func TestParams_ParseAndString(t *testing.T) {
	tests := []struct {
		name     string
		params   string
		wantErr  bool
		validate func(*testing.T, *Params)
	}{
		{
			name:    "Complete parameters",
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
			name:    "Partial parameters",
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
			name:    "Empty parameters",
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
			name:    "Boundary test - Maximum memory cost",
			params:  "t:3,m:4294967295,p:4,k:32",
			wantErr: false,
			validate: func(t *testing.T, p *Params) {
				if p.TimeCost != 3 {
					t.Errorf("TimeCost = %v, want %v", p.TimeCost, 3)
				}
				if p.MemoryCost != 4294967295 {
					t.Errorf("MemoryCost = %v, want %v", p.MemoryCost, 4294967295)
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
			name:    "Boundary test - Minimum memory cost",
			params:  "t:3,m:1,p:4,k:32",
			wantErr: false,
			validate: func(t *testing.T, p *Params) {
				if p.TimeCost != 3 {
					t.Errorf("TimeCost = %v, want %v", p.TimeCost, 3)
				}
				if p.MemoryCost != 1 {
					t.Errorf("MemoryCost = %v, want %v", p.MemoryCost, 1)
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
			name:    "Invalid parameter format",
			params:  "t:3,m:65536,p:4,k:32,invalid",
			wantErr: true,
		},
		{
			name:    "Invalid parameter value",
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

			// Test String method
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
			name: "Complete parameters",
			params: &Params{
				TimeCost:   3,
				MemoryCost: 65536,
				Threads:    4,
				KeyLength:  32,
			},
			want: "t:3,m:65536,p:4,k:32",
		},
		{
			name: "Partial parameters",
			params: &Params{
				TimeCost:   3,
				MemoryCost: 65536,
			},
			want: "t:3,m:65536",
		},
		{
			name:   "Zero value parameters",
			params: &Params{},
			want:   "",
		},
		{
			name: "Boundary test - Maximum threads",
			params: &Params{
				TimeCost:   3,
				MemoryCost: 65536,
				Threads:    255,
				KeyLength:  32,
			},
			want: "t:3,m:65536,p:255,k:32",
		},
		{
			name: "Boundary test - Minimum threads",
			params: &Params{
				TimeCost:   3,
				MemoryCost: 65536,
				Threads:    1,
				KeyLength:  32,
			},
			want: "t:3,m:65536,p:1,k:32",
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

func TestNewArgon2Crypto(t *testing.T) {
	tests := []struct {
		name    string
		config  *types.Config
		wantErr bool
	}{
		{
			name:    "Default config",
			config:  DefaultConfig(),
			wantErr: false,
		},
		{
			name: "Custom config",
			config: &types.Config{
				TimeCost:   3,
				MemoryCost: 64 * 1024,
				Threads:    4,
				SaltLength: 16,
				KeyLength:  32,
			},
			wantErr: false,
		},
		{
			name: "Invalid config - zero time cost",
			config: &types.Config{
				TimeCost:   0,
				MemoryCost: 64 * 1024,
				Threads:    4,
				SaltLength: 16,
				KeyLength:  32,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewArgon2Crypto(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewArgon2Crypto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCrypto_Hash(t *testing.T) {
	crypto, err := NewArgon2Crypto(DefaultConfig())
	if err != nil {
		t.Fatalf("Failed to create Argon2 crypto: %v", err)
	}

	hash, err := crypto.Hash("password")
	if err != nil {
		t.Errorf("Hash() error = %v", err)
	}
	if hash == "" {
		t.Error("Hash() returned empty string")
	}
}

func TestCrypto_HashWithSalt(t *testing.T) {
	crypto, err := NewArgon2Crypto(DefaultConfig())
	if err != nil {
		t.Fatalf("Failed to create Argon2 crypto: %v", err)
	}

	salt := "somesaltvalue"
	hash, err := crypto.HashWithSalt("password", salt)
	if err != nil {
		t.Errorf("HashWithSalt() error = %v", err)
	}
	if hash == "" {
		t.Error("HashWithSalt() returned empty string")
	}
}

func TestCrypto_Verify(t *testing.T) {
	crypto, err := NewArgon2Crypto(DefaultConfig())
	if err != nil {
		t.Fatalf("Failed to create Argon2 crypto: %v", err)
	}

	password := "password"
	hash, err := crypto.Hash(password)
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}

	err = crypto.Verify(hash, password)
	if err != nil {
		t.Errorf("Verify() error = %v", err)
	}

	// Test with wrong password
	wrongPassword := "wrongpassword"
	err = crypto.Verify(hash, wrongPassword)
	if err == nil {
		t.Error("Verify() should return error for wrong password")
	}
}

func TestCrypto_VerifyWithSalt(t *testing.T) {
	crypto, err := NewArgon2Crypto(DefaultConfig())
	if err != nil {
		t.Fatalf("Failed to create Argon2 crypto: %v", err)
	}
	password := "password"
	salt := "somesaltvalue"
	hash, err := crypto.HashWithSalt(password, salt)
	if err != nil {
		t.Fatalf("HashWithSalt() error = %v", err)
	}
	err = crypto.Verify(hash, password)
	if err != nil {
		t.Errorf("VerifyWithSalt() error = %v", err)
	}
	wrongPassword := "wrongpassword"
	err = crypto.Verify(hash, wrongPassword)
	if err == nil {
		t.Error("VerifyWithSalt() should return error for wrong password")
	}
}
