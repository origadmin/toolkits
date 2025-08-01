/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package interfaces

import "github.com/origadmin/toolkits/crypto/hash/types"

// ConfigValidator defines the interface for configuration validation
type ConfigValidator interface {
	// Validate validates the configuration
	Validate(*types.Config) error
}
