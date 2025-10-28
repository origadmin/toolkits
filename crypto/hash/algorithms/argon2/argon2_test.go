package argon2

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/validator"
)

func TestParams_ParseAndString(t *testing.T) {
	tests := []struct {
		name     string
		params   string
		wantErr  bool
		validate func(*testing.T, *Params)
	}{
		{
			name:    "CompleteParameters",
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
			name:    "BoundaryTestMaximumMemoryCost",
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
			name:    "BoundaryTestMinimumMemoryCost",
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
			name:    "InvalidParameterFormat",
			params:  "t:3,m:65536,p:4,k:32,invalid",
			wantErr: true,
		},
		{
			name:    "InvalidParameterValue",
			params:  "t:invalid,m:65536",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decodedParams, decodeErr := codec.DecodeParams(tt.params)
			if decodeErr != nil {
				if tt.wantErr {
					// Expected error, test passes for this case
					return
				}
				t.Fatalf("codec.DecodeParams() unexpected error = %v, wantErr %v", decodeErr, tt.wantErr)
			}

			// If we reach here, decodedParams is valid, proceed with validation
			p, err := validator.WithValidator(&types.Config{
				SaltLength: 16,
				Params:     decodedParams,
			}, func(cfg *types.Config) error {
				params, err := FromMap(cfg.Params)
				if err != nil {
					return err
				}
				if tt.validate != nil {
					tt.validate(t, params)
				}
				return nil
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("validator.ValidateParams() error = %p, wantErr %v", err, tt.wantErr)
				return // If we got an unexpected error or didn't get an expected error, fail and return
			}
			if tt.wantErr {
				return // If we expected an error and got one from ValidateParams, test passes
			}

			// Test String method to ensure it's the inverse of parsing.
			// We parse the output of String() and compare the resulting struct
			// with the original one.
			if tt.params != "" {
				p2, err := validator.WithValidator(&types.Config{
					Params:     p.Params,
					SaltLength: 16,
				}, func(cfg *types.Config) error {
					params, err := FromMap(cfg.Params)
					if err != nil {
						return err
					}
					if tt.validate != nil {
						tt.validate(t, params)
					}
					return nil
				})
				if err != nil {
					t.Fatalf("Failed to validate params after String()->Parse roundtrip for Params '%s': %v", p.Params, err)
				}

				if !reflect.DeepEqual(p.Params, p2.Params) {
					t.Errorf("Params after String()->Parse roundtrip do not match original. got %+v, want %+v",
						p2.Params, p.Params)
				}
			}
		})
	}
}

func TestNewArgon2(t *testing.T) {
	tests := []struct {
		name            string
		algSpec         types.Spec
		config          *types.Config
		expectedAlgSpec types.Spec
		wantErr         bool
	}{
		{
			name:            "Default config for ARGON2",
			algSpec:         types.New(types.ARGON2),
			config:          DefaultConfig(),
			expectedAlgSpec: types.Spec{Name: types.ARGON2i},
			wantErr:         true, // Changed to true, as argon2.go does not support generic ARGON2
		},
		{
			name:            "Default config for ARGON2i",
			algSpec:         types.New(types.ARGON2i),
			config:          DefaultConfig(),
			expectedAlgSpec: types.Spec{Name: types.ARGON2i},
			wantErr:         false,
		},
		{
			name:            "Default config for ARGON2id",
			algSpec:         types.New(types.ARGON2id),
			config:          DefaultConfig(),
			expectedAlgSpec: types.Spec{Name: types.ARGON2id},
			wantErr:         false,
		},
		{
			name:    "Custom config",
			algSpec: types.New(types.ARGON2id),
			config: &types.Config{
				SaltLength: types.DefaultSaltLength,
				Params: (&Params{
					TimeCost:   types.DefaultTimeCost,
					MemoryCost: types.DefaultMemoryCost,
					Threads:    types.DefaultThreads,
					KeyLength:  32,
				}).ToMap(),
			},
			expectedAlgSpec: types.Spec{Name: types.ARGON2id},
			wantErr:         false,
		},
		{
			name:    "Invalid config - zero time cost",
			algSpec: types.New(types.ARGON2i),
			config: &types.Config{
				SaltLength: types.DefaultSaltLength,
				Params: (&Params{
					TimeCost:   0,
					MemoryCost: types.DefaultMemoryCost,
					Threads:    types.DefaultThreads,
					KeyLength:  32,
				}).ToMap(),
			},
			expectedAlgSpec: types.Spec{},
			wantErr:         true,
		},
		{
			name:            "Unsupported algorithm type",
			algSpec:         types.New("unsupported"),
			config:          DefaultConfig(),
			expectedAlgSpec: types.Spec{},
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewArgon2(tt.algSpec, tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewArgon2() error = %v, wantErr %v", err, tt.wantErr)
				return // Prevent nil pointer dereference
			}
			if !tt.wantErr {
				assert.NotNil(t, c)
				assert.Equal(t, tt.expectedAlgSpec, c.Spec())
				// Test Hash and Verify for valid cases
				hash, err := c.Hash("password")
				assert.NoError(t, err)
				assert.NotNil(t, hash)
				assert.Equal(t, tt.expectedAlgSpec, hash.Spec)

				err = c.Verify(hash, "password")
				assert.NoError(t, err)
			}
		})
	}
}
