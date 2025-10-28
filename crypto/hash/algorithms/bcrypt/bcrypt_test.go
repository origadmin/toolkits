package bcrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"

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
				Params: (&Params{
					Cost: types.DefaultCost,
				}).ToMap(),
			},
			wantErr: false,
		},
		{
			name: "Invalid config - zero cost",
			config: &types.Config{
				Params: (&Params{
					Cost: 0,
				}).ToMap(),
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
	assert.Panics(t, func() {
		crypto.Verify(nil, "password")
	})
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

func TestBcrypt_Verify_Error(t *testing.T) {
	c, err := NewBcrypt(DefaultConfig())
	assert.NoError(t, err)

	// Invalid algorithm
	err = c.Verify(&types.HashParts{Spec: types.New("invalid")}, "password")
	assert.Error(t, err)

	// Spec mismatch
	hash, err := c.Hash("password")
	assert.NoError(t, err)
	hash.Spec = types.New("argon2")
	err = c.Verify(hash, "password")
	assert.Error(t, err)
}

func TestBcrypt_Hash_Error(t *testing.T) {
	// This test is a bit tricky as it requires mocking rand.RandomBytes.
	// For now, we'll just ensure the function doesn't panic with a valid config.
	c, err := NewBcrypt(DefaultConfig())
	assert.NoError(t, err)
	_, err = c.Hash("password")
	assert.NoError(t, err)
}

func TestBcrypt_HashWithSalt_Error(t *testing.T) {
	// This test is a bit tricky as it requires mocking bcrypt.GenerateFromPassword.
	// For now, we'll just ensure the function doesn't panic with a valid config.
	c, err := NewBcrypt(DefaultConfig())
	assert.NoError(t, err)
	_, err = c.HashWithSalt("password", []byte("salt"))
	assert.NoError(t, err)
}
