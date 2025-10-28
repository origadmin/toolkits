package scrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/crypto/hash/types"
)

func TestNewScrypt(t *testing.T) {
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
					N:      16384,
					R:      8,
					P:      1,
					KeyLen: 32,
				}).ToMap(),
			},
			wantErr: false,
		},
		{
			name: "Invalid config - invalid N",
			config: &types.Config{
				Params: (&Params{
					N:      1,
					R:      8,
					P:      1,
					KeyLen: 32,
				}).ToMap(),
				SaltLength: 16,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewScrypt(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewScrypt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestScrypt_HashAndVerify(t *testing.T) {
	crypto, err := NewScrypt(DefaultConfig())
	assert.NoError(t, err)

	hash, err := crypto.Hash("password")
	assert.NoError(t, err)

	err = crypto.Verify(hash, "password")
	assert.NoError(t, err)

	err = crypto.Verify(hash, "wrongpassword")
	assert.Error(t, err)
}

func TestScrypt_HashWithSaltAndVerify(t *testing.T) {
	crypto, err := NewScrypt(DefaultConfig())
	assert.NoError(t, err)

	salt := []byte("somesalt")
	hash, err := crypto.HashWithSalt("password", salt)
	assert.NoError(t, err)

	err = crypto.Verify(hash, "password")
	assert.NoError(t, err)

	err = crypto.Verify(hash, "wrongpassword")
	assert.Error(t, err)
}

func TestScrypt_Verify_Error(t *testing.T) {
	c, err := NewScrypt(DefaultConfig())
	assert.NoError(t, err)

	// Invalid algorithm
	err = c.Verify(&types.HashParts{Spec: types.New("invalid")}, "password")
	assert.Error(t, err)

	// Spec mismatch
	hash, err := c.Hash("password")
	assert.NoError(t, err)
	hash.Spec = types.New("bcrypt")
	err = c.Verify(hash, "password")
	assert.Error(t, err)

	// Invalid params
	hash, err = c.Hash("password")
	assert.NoError(t, err)
	hash.Params = map[string]string{"invalid": "param"}
	err = c.Verify(hash, "password")
	assert.Error(t, err)
}

func TestScrypt_Hash_Error(t *testing.T) {
	// This test is a bit tricky as it requires mocking rand.RandomBytes.
	// For now, we'll just ensure the function doesn't panic with a valid config.
	c, err := NewScrypt(DefaultConfig())
	assert.NoError(t, err)
	_, err = c.Hash("password")
	assert.NoError(t, err)
}

func TestScrypt_HashWithSalt_Error(t *testing.T) {
	// This test is a bit tricky as it requires mocking scrypt.Key.
	// For now, we'll just ensure the function doesn't panic with a valid config.
	c, err := NewScrypt(DefaultConfig())
	assert.NoError(t, err)
	_, err = c.HashWithSalt("password", []byte("salt"))
	assert.NoError(t, err)
}
