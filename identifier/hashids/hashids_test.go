/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package hashids_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier/hashids"
)

func TestNew(t *testing.T) {
	// Test successful creation
	cfg := hashids.Config{
		Salt:      "my-secret-salt",
		MinLength: 8,
	}
	h, err := hashids.New(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, h)

	// Test creation failure (empty salt)
	_, err = hashids.New(hashids.Config{})
	assert.Error(t, err, "Expected an error when creating with an empty salt")
}

func TestEncodeDecode(t *testing.T) {
	cfg := hashids.Config{
		Salt:      "this is a test",
		MinLength: 10,
	}
	h, _ := hashids.New(cfg)

	// Test with a single number
	numbers := []int64{12345}
	hash, err := h.Encode(numbers...)
	assert.NoError(t, err)
	assert.True(t, len(hash) >= 10)

	decoded, err := h.Decode(hash)
	assert.NoError(t, err)
	assert.Equal(t, numbers, decoded)

	// Test with multiple numbers
	multiNumbers := []int64{6, 8, 3, 9, 4, 2}
	multiHash, err := h.Encode(multiNumbers...)
	assert.NoError(t, err)

	multiDecoded, err := h.Decode(multiHash)
	assert.NoError(t, err)
	assert.Equal(t, multiNumbers, multiDecoded)
}

func TestEncodeDecodeWithCustomAlphabet(t *testing.T) {
	cfg := hashids.Config{
		Salt:     "custom alphabet test",
		Alphabet: "abcdef1234567890", // Must be at least 16 chars
	}
	h, _ := hashids.New(cfg)

	numbers := []int64{1, 2, 3}
	hash, err := h.Encode(numbers...)
	assert.NoError(t, err)

	// Check that the hash only contains characters from the custom alphabet
	for _, r := range hash {
		assert.Contains(t, cfg.Alphabet, string(r))
	}

	decoded, err := h.Decode(hash)
	assert.NoError(t, err)
	assert.Equal(t, numbers, decoded)
}

func TestDecodeInvalidHash(t *testing.T) {
	cfg := hashids.Config{Salt: "test"}
	h, _ := hashids.New(cfg)

	t.Run("InvalidHash", func(t *testing.T) {
		// An invalid hash string should return an error.
		decoded, err := h.Decode("not a valid hash")
		assert.Error(t, err)
		assert.Empty(t, decoded)
	})

	t.Run("EmptyHash", func(t *testing.T) {
		// An empty hash string should decode to an empty slice without error.
		decoded, err := h.Decode("")
		assert.NoError(t, err)
		assert.Empty(t, decoded)
	})
}
