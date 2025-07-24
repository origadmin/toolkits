package bcrypt

import (
	"testing"

	codecPkg "github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

var codec = codecPkg.NewCodec(types.TypeBcrypt)

func TestNewBcryptCrypto(t *testing.T) {
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
				SaltLength: 16,
				ParamConfig: (&Params{
					Cost: constants.DefaultCost,
				}).String(),
			},
			wantErr: false,
		},
		{
			name: "Invalid config - zero cost",
			config: &types.Config{
				ParamConfig: (&Params{
					Cost: 0,
				}).String(),
				SaltLength: 16,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewBcryptCrypto(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBcryptCrypto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCrypto_Hash(t *testing.T) {
	crypto, err := NewBcryptCrypto(DefaultConfig())
	if err != nil {
		t.Fatalf("Failed to create Bcrypt crypto: %v", err)
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
	crypto, err := NewBcryptCrypto(DefaultConfig())
	if err != nil {
		t.Fatalf("Failed to create Bcrypt crypto: %v", err)
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
	crypto, err := NewBcryptCrypto(DefaultConfig())
	if err != nil {
		t.Fatalf("Failed to create Bcrypt crypto: %v", err)
	}

	password := "password"
	hash, err := crypto.Hash(password)
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}
	decoded, err := codec.Decode(hash)
	if err != nil {
		t.Errorf("Decode() error = %v", err)

	}
	err = crypto.Verify(decoded, password)
	if err != nil {
		t.Errorf("Verify() error = %v", err)
	}

	// Test with wrong password
	wrongPassword := "wrongpassword"
	err = crypto.Verify(decoded, wrongPassword)
	if err == nil {
		t.Error("Verify() should return error for wrong password")
	}
}

func TestCrypto_VerifyWithSalt(t *testing.T) {
	crypto, err := NewBcryptCrypto(DefaultConfig())
	if err != nil {
		t.Fatalf("Failed to create Bcrypt crypto: %v", err)
	}
	password := "password"
	salt := "somesaltvalue"
	hash, err := crypto.HashWithSalt(password, salt)
	if err != nil {
		t.Fatalf("HashWithSalt() error = %v", err)
	}
	decoded, err := codec.Decode(hash)
	if err != nil {
		t.Errorf("Decode() error = %v", err)

	}
	err = crypto.Verify(decoded, password)
	if err != nil {
		t.Errorf("VerifyWithSalt() error = %v", err)
	}
	wrongPassword := "wrongpassword"
	err = crypto.Verify(decoded, wrongPassword)
	if err == nil {
		t.Error("VerifyWithSalt() should return error for wrong password")
	}
}
