/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package types

// ScryptConfig represents the configuration for Scrypt algorithm
type ScryptConfig struct {
	N      int `env:"HASH_SCRYPT_N"`
	R      int `env:"HASH_SCRYPT_R"`
	P      int `env:"HASH_SCRYPT_P"`
	KeyLen int `env:"HASH_SCRYPT_KEYLEN"`
}

// Config represents the configuration for hash algorithms
type Config struct {
	Algorithm  Type   `env:"HASH_ALGORITHM"`
	TimeCost   uint32 `env:"HASH_TIMECOST"`
	MemoryCost uint32 `env:"HASH_MEMORYCOST"`
	Threads    uint8  `env:"HASH_THREADS"`
	SaltLength int    `env:"HASH_SALTLENGTH"`
	Cost       int    `env:"HASH_COST"` // Cost parameter (for bcrypt)
	Salt       string `env:"HASH_SALT"` // Salt for HMAC
	Scrypt     ScryptConfig
}

// ConfigOption is a function that modifies a Config
type ConfigOption func(*Config)

// WithAlgorithm sets the algorithm type
func WithAlgorithm(t Type) ConfigOption {
	return func(cfg *Config) {
		cfg.Algorithm = t
	}
}

// WithTimeCost sets the time cost
func WithTimeCost(cost uint32) ConfigOption {
	return func(cfg *Config) {
		cfg.TimeCost = cost
	}
}

// WithMemoryCost sets the memory cost
func WithMemoryCost(cost uint32) ConfigOption {
	return func(cfg *Config) {
		cfg.MemoryCost = cost
	}
}

// WithThreads sets the number of threads
func WithThreads(threads uint8) ConfigOption {
	return func(cfg *Config) {
		cfg.Threads = threads
	}
}

// WithSaltLength sets the salt length
func WithSaltLength(length int) ConfigOption {
	return func(cfg *Config) {
		cfg.SaltLength = length
	}
}

// WithCost sets the cost parameter
func WithCost(cost int) ConfigOption {
	return func(cfg *Config) {
		cfg.Cost = cost
	}
}

// WithSalt sets the salt
func WithSalt(salt string) ConfigOption {
	return func(cfg *Config) {
		cfg.Salt = salt
	}
}

// WithScryptConfig sets the scrypt configuration
func WithScryptConfig(scrypt ScryptConfig) ConfigOption {
	return func(cfg *Config) {
		cfg.Scrypt = scrypt
	}
}
