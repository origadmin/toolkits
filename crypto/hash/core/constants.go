/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package core

const (
	// ENV environment variable name
	ENV = "ORIGADMIN_HASH_TYPE"
	// DefaultType default hash type
	DefaultType = "argon2"
	// DefaultVersion default hash version
	DefaultVersion = "v1"
	// DefaultSaltLength default salt length
	DefaultSaltLength = 16
	// DefaultTimeCost default time cost for Argon2
	DefaultTimeCost = 3
	// DefaultMemoryCost default memory cost for Argon2
	DefaultMemoryCost = 64 * 1024 // 64MB
	// DefaultThreads default number of threads for Argon2
	DefaultThreads = 4
	// DefaultCost default cost for bcrypt
	DefaultCost = 10

	ParamSeparator      = ","
	ParamValueSeparator = ":"
)
