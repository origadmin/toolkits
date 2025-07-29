package bcrypt

import (
	"testing"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

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
			_, err := NewBcrypt(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBcrypt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCrypto_Hash(t *testing.T) {
	crypto, err := NewBcrypt(DefaultConfig())
	if err != nil {
		t.Fatalf("Failed to create Bcrypt crypto: %v", err)
	}

	hashParts, err := crypto.Hash("password")
	if err != nil {
		t.Errorf("Hash() error = %v", err)
	}
	err = crypto.Verify(hashParts, "password")
	if err != nil {
		t.Errorf("Verify() error = %v", err)
	}
}

func TestCrypto_HashWithSalt(t *testing.T) {
	crypto, err := NewBcrypt(DefaultConfig())
	if err != nil {
		t.Fatalf("Failed to create Bcrypt crypto: %v", err)
	}

	salt := []byte("somesaltvalue")
	hashParts, err := crypto.HashWithSalt("password", salt)
	if err != nil {
		t.Errorf("HashWithSalt() error = %v", err)
	}
	err = crypto.Verify(hashParts, "password")
	if err != nil {
		t.Errorf("Verify() error = %v", err)
	}
	// Test with wrong password
	wrongPassword := "wrongpassword"
	err = crypto.Verify(hashParts, wrongPassword)
	if err == nil {
		t.Error("Verify() should return error for wrong password")
	}
	// Test with empty password
	err = crypto.Verify(hashParts, "")
	if err == nil {
		t.Error("Verify() should return error for empty password")
	}
	// Test with nil hash
	err = crypto.Verify(nil, "password")
	if err == nil {
		t.Error("Verify() should return error for nil hash")
	}
}

func TestCrypto_Verify(t *testing.T) {
	crypto, err := NewBcrypt(DefaultConfig())
	if err != nil {
		t.Fatalf("Failed to create Bcrypt crypto: %v", err)
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
	crypto, err := NewBcrypt(DefaultConfig())
	if err != nil {
		t.Fatalf("Failed to create Bcrypt crypto: %v", err)
	}
	password := "password"
	salt := []byte("somesaltvalue")
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
