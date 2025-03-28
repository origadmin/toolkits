/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package types

// ScryptConfig represents the configuration for Scrypt algorithm
type ScryptConfig struct {
	N int `env:"HASH_SCRYPT_N"`
	R int `env:"HASH_SCRYPT_R"`
	P int `env:"HASH_SCRYPT_P"`
}

// Config represents the configuration for hash algorithms
type Config struct {
	TimeCost   uint32 `env:"HASH_TIMECOST"`
	MemoryCost uint32 `env:"HASH_MEMORYCOST"`
	Threads    uint8  `env:"HASH_THREADS"`
	SaltLength int    `env:"HASH_SALTLENGTH"`
	KeyLength  uint32 `env:"HASH_KEYLENGTH"`
	Cost       int    `env:"HASH_COST"` // Cost parameter (for bcrypt)
	Salt       string `env:"HASH_SALT"` // Salt for HMAC
	Iterations int    `env:"HASH_ITERATIONS"`
	HashType   string `env:"HASH_TYPE"` // used for pbkdf2
	Scrypt     ScryptConfig
}

// ConfigOption is a function that modifies a Config
type ConfigOption func(*Config)

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

// WithIterations sets the number of iterations for pbkdf2
func WithIterations(iterations int) ConfigOption {
	return func(cfg *Config) {
		cfg.Iterations = iterations
	}
}

// WithHashType sets the hash type for pbkdf2
func WithHashType(hashType string) ConfigOption {
	return func(cfg *Config) {
		cfg.HashType = hashType
	}
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		TimeCost:   3,        // Default time cost
		MemoryCost: 65536,    // Default memory cost (64MB)
		Threads:    4,        // Default number of threads
		SaltLength: 16,       // Default salt length
		KeyLength:  32,       // Default key length
		Cost:       10,       // Default cost parameter
		Salt:       "",       // Default salt
		Iterations: 100000,   // Default iterations for pbkdf2
		HashType:   "sha256", // Default hash type for pbkdf2
		Scrypt: ScryptConfig{
			N: 16384, // Default N value
			R: 8,     // Default R value
			P: 1,     // Default P value
		},
	}
}
