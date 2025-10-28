/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package examples

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/bcrypt"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/scheme"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Example demonstrating basic hashing and verification.
func Example_basic() {
	password := "my-secret-password"

	// Create a new crypto instance with a standard algorithm (e.g., SHA256)
	c, err := hash.NewCrypto(types.SHA256)
	if err != nil {
		log.Fatalf("Failed to create crypto: %v", err)
	}

	// Hash the password. The result is an encoded string containing all necessary info.
	hashed, err := c.Hash(password)
	if err != nil {
		log.Fatalf("Hashing failed: %v", err)
	}

	// Verification should succeed with the correct password.
	if err := c.Verify(hashed, password); err != nil {
		fmt.Println("Verification failed with correct password.")
	} else {
		fmt.Println("Verification successful!")
	}

	// Verification should fail with an incorrect password.
	if err := c.Verify(hashed, "wrong-password"); err != nil {
		fmt.Println("Verification correctly failed for wrong password.")
	}

	// Output:
	// Verification successful!
	// Verification correctly failed for wrong password.
}

// Example demonstrating how to use an algorithm with specific options, like bcrypt cost.
func Example_withOptions() {
	password := "another-secure-password"

	// Create a bcrypt hasher with a custom cost.
	// Note: Higher cost is more secure but slower.
	c, err := hash.NewCrypto(types.BCRYPT, bcrypt.WithCost(10)) // Using a cost of 10 for the example
	if err != nil {
		log.Fatalf("Failed to create bcrypt crypto: %v", err)
	}

	// Hash and verify
	hashed, err := c.Hash(password)
	if err != nil {
		log.Fatalf("Hashing with bcrypt failed: %v", err)
	}

	if err := c.Verify(hashed, password); err == nil {
		fmt.Println("Bcrypt verification successful!")
	}

	// Output: Bcrypt verification successful!
}

// simpleCustomHasher is a dummy implementation of the scheme.Scheme interface for demonstration.
type simpleCustomHasher struct {
	spec types.Spec
}

func (h *simpleCustomHasher) Spec() types.Spec {
	return h.spec
}

func (h *simpleCustomHasher) Hash(password string) (*types.HashParts, error) {
	salt := make([]byte, 8) // Using a smaller salt for the example
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

// Example demonstrating how to register and use a custom hashing algorithm.
func Example_customAlgorithm() {
	// Register the custom hasher implementation
	// Create a spec that will be parsed into Name="simple" and Underlying="custom".
	// This ensures the factory is registered under the name "simple".
	spec := types.New("simple", "custom")
	customAlgName := "simple-custom"
	customAlgAlias := "custom"
	hash.Register(&simpleCustomFactory{}, spec, customAlgName, customAlgAlias)

	// Create a crypto instance using the custom algorithm's name
	c, err := hash.NewCrypto(customAlgName)
	if err != nil {
		log.Fatalf("Failed to create custom crypto: %v", err)
	}

	password := "custom-password"
	hashed, _ := c.Hash(password)

	if err := c.Verify(hashed, password); err == nil {
		fmt.Println("Custom algorithm verification successful!")
	} else {
		fmt.Printf("Custom algorithm verification failed: %v", err)
	}

	// Now, create an instance using its alias
	cFromAlias, err := hash.NewCrypto(customAlgAlias)
	if err != nil {
		log.Fatalf("Failed to create custom crypto from alias: %v", err)
	}

	if err := cFromAlias.Verify(hashed, password); err == nil {
		fmt.Println("Custom algorithm (alias) verification successful!")
	}

	// Output:
	// Custom algorithm verification successful!
	// Custom algorithm (alias) verification successful!
}
