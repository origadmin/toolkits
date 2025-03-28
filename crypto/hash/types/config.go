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

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		TimeCost:   3,     // 默认时间成本
		MemoryCost: 65536, // 默认内存成本 (64MB)
		Threads:    4,     // 默认线程数
		SaltLength: 16,    // 默认盐长度
		KeyLength:  32,    // 默认密钥长度
	}
}
