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

// Example_basic demonstrates the most common use case: creating a hash with a
// standard algorithm and verifying it.
func Example_basic() {
	password := "my-secret-password"

	// 1. Create a new crypto instance with a standard algorithm (e.g., SHA256).
	//    The returned `Crypto` instance is configured to create new hashes using SHA256.
	c, err := hash.NewCrypto(types.SHA256)
	if err != nil {
		log.Fatalf("Failed to create crypto: %v", err)
	}

	// 2. Hash the password. The result is a single, encoded string that contains the
	//    algorithm name, parameters, salt, and the hash itself.
	hashed, err := c.Hash(password)
	if err != nil {
		log.Fatalf("Hashing failed: %v", err)
	}

	// 3. Verify the password. The `Verify` method automatically detects the algorithm
	//    from the hashed string and uses the correct logic to compare.
	if err := c.Verify(hashed, password); err != nil {
		fmt.Println("Verification failed with correct password.")
	} else {
		fmt.Println("Verification successful!")
	}

	// 4. Verification with an incorrect password should always return an error.
	if err := c.Verify(hashed, "wrong-password"); err != nil {
		fmt.Println("Verification correctly failed for wrong password.")
	}

	// Output:
	// Verification successful!
	// Verification correctly failed for wrong password.
}

// Example_withOptions demonstrates how to create a crypto instance with algorithm-specific
// options, such as setting the cost for bcrypt.
func Example_withOptions() {
	password := "another-secure-password"

	// 1. When creating the instance, pass in option functions from the specific algorithm's package.
	//    Here, we use `bcrypt.WithCost` to set a non-default cost factor.
	//    Note: A higher cost is more secure but significantly slower.
	c, err := hash.NewCrypto(types.BCRYPT, bcrypt.WithCost(10)) // Using a cost of 10 for the example
	if err != nil {
		log.Fatalf("Failed to create bcrypt crypto: %v", err)
	}

	// 2. Hash and verify as usual. The cost parameter is automatically encoded into the hash string.
	hashed, err := c.Hash(password)
	if err != nil {
		log.Fatalf("Hashing with bcrypt failed: %v", err)
	}

	if err := c.Verify(hashed, password); err == nil {
		fmt.Println("Bcrypt verification successful!")
	}

	// Output: Bcrypt verification successful!
}

// --- Custom Algorithm Demonstration ---

// simpleCustomHasher is a dummy implementation of the scheme.Scheme interface for demonstration.
// It "hashes" by simply reversing the password string.
// DO NOT USE THIS IN PRODUCTION.
type simpleCustomHasher struct {
	spec types.Spec
}

func (h *simpleCustomHasher) Spec() types.Spec {
	return h.spec
}

func (h *simpleCustomHasher) Hash(password string) (*types.HashParts, error) {
	salt := make([]byte, 8) // A real algorithm would use a cryptographically secure salt.
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return h.HashWithSalt(password, salt)
}

func (h *simpleCustomHasher) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	// The actual hashing logic is trivial for this example.
	hashBytes := h.reverse(password)

	// Use the NewHashParts constructor to assemble the final parts.
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

// simpleCustomFactory creates instances of our simpleCustomHasher.
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

// Example_customAlgorithm demonstrates how to register and use a custom hashing algorithm.
func Example_customAlgorithm() {
	// 1. Define the algorithm's canonical spec and any aliases.
	//    Because our name "simple-custom" contains a hyphen, we must register it as a
	//    composite algorithm to work with the default parser.
	spec := types.New("simple", "custom")
	customAlgName := "simple-custom"
	customAlgAlias := "custom"

	// 2. Register your factory with its spec and aliases. This makes the algorithm
	//    available to the entire application through `hash.NewCrypto`.
	hash.Register(&simpleCustomFactory{}, spec, customAlgName, customAlgAlias)

	// 3. Now, you can create a crypto instance using your custom algorithm's name.
	c, err := hash.NewCrypto(customAlgName)
	if err != nil {
		log.Fatalf("Failed to create custom crypto: %v", err)
	}

	password := "custom-password"
	hashed, _ := c.Hash(password)

	// 4. Verification works seamlessly, just like with a built-in algorithm.
	if err := c.Verify(hashed, password); err == nil {
		fmt.Println("Custom algorithm verification successful!")
	}

	// 5. You can also use the registered alias.
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
