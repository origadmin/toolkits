/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"fmt"

	"github.com/goexts/generic/configure"

	"github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

type Crypto interface {
	Spec() types.Spec
	Hash(password string) (string, error)
	HashWithSalt(password string, salt []byte) (string, error)
	Verify(hashed, password string) error
}

type crypto struct {
	codec codec.Codec
	// algImpl is the cryptographic implementation used for hashing, wrapped with cachedVerifier
	algImpl interfaces.Cryptographic
}

func (c *crypto) Spec() types.Spec {
	return c.algImpl.Spec()
}

func (c *crypto) Hash(password string) (string, error) {
	hashParts, err := c.algImpl.Hash(password)
	if err != nil {
		return "", err
	}
	return c.codec.Encode(hashParts)
}

func (c *crypto) HashWithSalt(password string, salt []byte) (string, error) {
	hashParts, err := c.algImpl.HashWithSalt(password, salt)
	if err != nil {
		return "", err
	}
	return c.codec.Encode(hashParts)
}

func (c *crypto) Verify(hashed, password string) error {
	// Check for nil or empty hashed string early
	if hashed == "" {
		return errors.ErrInvalidHash
	}

	// Decode the hash value to get algorithm type and parts
	parts, err := c.codec.Decode(hashed)
	if err != nil {
		return err
	}

	// Perform nil checks for HashParts directly here
	if parts == nil || parts.Hash == nil || parts.Salt == nil {
		return errors.ErrInvalidHashParts
	}

	// Get algorithm instance from global factory based on the decoded algorithm
	factory := getFactory()
	cryptographic, err := factory.create(parts.Algorithm, WithEncodedParams(parts.Params,
		codec.EncodeParams)) // Pass types.Spec directly
	if err != nil {
		return err
	}

	// Wrap the created algorithm with a cached verifier for performance
	cachedAlg := NewCachedVerifier(cryptographic)

	algEntry, exists := algorithmMap[parts.Algorithm.Name]
	if !exists {
		return fmt.Errorf("unsupported algorithm: %s", parts.Algorithm.String())
	}
	parts.Algorithm, err = algEntry.resolver.ResolveSpec(parts.Algorithm) // Ensure the resolver is called
	if err != nil {
		return err
	}

	return cachedAlg.Verify(parts, password)
}

// NewCrypto creates a new cryptographic instance
func NewCrypto(algName string, opts ...Option) (Crypto, error) {
	// 1. Parse the algorithm name string into a structured Spec
	algSpec, err := types.Parse(algName)
	if err != nil {
		return nil, err
	}

	// 2. Look up the algorithm entry to get its specific resolver
	algEntry, exists := algorithmMap[algSpec.Name]
	if !exists {
		return nil, fmt.Errorf("unsupported algorithm: %s", algSpec.String())
	}

	// 3. Resolve the parsed Spec to its canonical form using the algorithm's specific resolver
	resolvedAlgSpec, err := algEntry.resolver.ResolveSpec(algSpec) // Use algEntry.resolver
	if err != nil {
		return nil, fmt.Errorf("failed to resolve algorithm type %s: %w", algSpec.String(), err)
	}

	// 4. Apply options to default config and create the cryptographic instance
	cfg := configure.Apply(algEntry.defaultConfig(), opts)
	cryptographic, err := algEntry.creator(resolvedAlgSpec, cfg) // Pass the resolved type
	if err != nil {
		return nil, err
	}

	// Wrap the cryptographic implementation with a cached verifier for hashing operations
	finalAlg := NewCachedVerifier(cryptographic)

	// Create cryptographic instance
	return &crypto{
			algImpl: finalAlg,
			codec:   codec.NewCodec(),
		},
		nil
}

// RegisterAlgorithm register a new hash algorithm that uses defaultSpecResolver by default
func RegisterAlgorithm(algSpec types.Spec, creator interfaces.AlgorithmCreator, defaultConfig interfaces.AlgorithmConfig) {
	algorithmMap[algSpec.Name] = algorithm{
		algSpec:       algSpec,
		creator:       creator,
		defaultConfig: defaultConfig,
		resolver:      defaultSpecResolver,
	}
}

// RegisterAlgorithmWithResolver register a new hash algorithm and specify a custom SpecResolver
func RegisterAlgorithmWithResolver(algSpec types.Spec, creator interfaces.AlgorithmCreator, defaultConfig interfaces.AlgorithmConfig, resolver interfaces.SpecResolver) {
	if resolver == nil {
		resolver = defaultSpecResolver
	}

	algorithmMap[algSpec.Name] = algorithm{
		algSpec:       algSpec,
		creator:       creator,
		defaultConfig: defaultConfig,
		resolver:      resolver,
	}
}

// AlgorithmMap returns all registered algorithm mapping tables (read-only access)
func AlgorithmMap() map[string]algorithm {
	// 返回一个副本以防止外部修改
	result := make(map[string]algorithm, len(algorithmMap))
	for k, v := range algorithmMap {
		result[k] = v
	}
	return result
}
