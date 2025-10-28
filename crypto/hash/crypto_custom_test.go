/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package hash

import (
	"crypto/rand"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/origadmin/toolkits/crypto/hash/algorithms/bcrypt"
	"github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/scheme"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// simpleCustomHasher is a dummy implementation of the scheme.Scheme interface for testing purposes.
type simpleCustomHasher struct {
	spec types.Spec
}

func (h *simpleCustomHasher) Spec() types.Spec {
	return h.spec
}

func (h *simpleCustomHasher) Hash(password string) (*types.HashParts, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return h.HashWithSalt(password, salt)
}

func (h *simpleCustomHasher) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	// This is NOT a secure hash. For demonstration purposes only.
	hashBytes := h.reverse(password)
	parts := types.NewHashParts(h.Spec(), []byte(hashBytes), salt, nil)
	return parts, nil
}

func (h *simpleCustomHasher) Verify(parts *types.HashParts, password string) error {
	expectedHash := h.reverse(password)
	if string(parts.Hash) != expectedHash {
		return errors.ErrPasswordNotMatch
	}
	return nil
}

func (h *simpleCustomHasher) reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// simpleCustomFactory is a factory for creating simpleCustomHasher instances.
type simpleCustomFactory struct{}

func (f *simpleCustomFactory) Create(spec types.Spec, _ *types.Config) (scheme.Scheme, error) {
	return &simpleCustomHasher{spec: spec}, nil
}

func (f *simpleCustomFactory) Config() *types.Config {
	return types.DefaultConfig()
}

func (f *simpleCustomFactory) ResolveSpec(spec types.Spec) (types.Spec, error) {
	return spec, nil
}

func TestCustomAlgorithmRegistration(t *testing.T) {
	customAlgName := "simple-custom"
	customAlgAlias := "custom"

	// 1. Register the new algorithm
	factory := &simpleCustomFactory{}
	// Create a spec that will be parsed into Name="simple" and Underlying="custom".
	// This ensures the factory is registered under the name "simple".
	spec := types.New("simple", "custom")
	Register(factory, spec, customAlgName, customAlgAlias)

	// 2. Create a crypto instance using the full name
	crypto, err := NewCrypto(customAlgName)
	require.NoError(t, err, "Failed to create crypto with custom algorithm")

	// 3. Test hashing and verification
	password := "password123"
	hashedString, err := crypto.Hash(password)
	require.NoError(t, err, "Hashing with custom algorithm failed")

	// The hashed string should now be in the encoded format, e.g., "$simple-custom$..."
	assert.True(t, strings.HasPrefix(hashedString, "$"+customAlgName), "Hash should have the correct prefix")

	err = crypto.Verify(hashedString, password)
	assert.NoError(t, err, "Verification with correct password failed")

	err = crypto.Verify(hashedString, "wrongpassword")
	assert.Error(t, err, "Verification with incorrect password should fail")
}

func TestBcryptCostOption(t *testing.T) {
	password := "my-secret-password"
	cost := 12
	costStr := strconv.Itoa(cost)

	// 1. Create a crypto instance with a specific bcrypt cost.
	crypto, err := NewCrypto(types.BCRYPT, bcrypt.WithCost(cost))
	require.NoError(t, err)

	// 2. Hash the password.
	hashed, err := crypto.Hash(password)
	require.NoError(t, err)

	// 3. Decode the resulting hash string using the library's own codec.
	// This is the correct way to test, by using the public API to verify internal state.
	decoder := codec.NewCodec()
	parts, err := decoder.Decode(hashed)
	require.NoError(t, err, "Failed to decode the generated hash")
	t.Logf("Decoded parts: %+v from hash: %s", parts, hashed)
	// 4. Verify that the cost parameter in the decoded parts matches the one we set.
	require.NotNil(t, parts.Params, "Params map should not be nil")
	extractedCost, ok := parts.Params["c"]
	require.True(t, ok, "Cost parameter 'c' not found in decoded hash params")
	assert.Equal(t, costStr, extractedCost, "The cost option was not applied or encoded correctly")

	// 5. Finally, ensure the password still verifies correctly.
	err = crypto.Verify(hashed, password)
	require.NoError(t, err, "Password verification failed with the generated hash")
}
