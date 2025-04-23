/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package helpers implements the functions, types, and interfaces for the module.
package helpers

import (
	"time"
)

// Int64ToDuration converts an integer value to a time.Duration type.
func Int64ToDuration(seconds int64) time.Duration {
	return time.Duration(seconds) * 1e6
}

func DurationToInt64(duration time.Duration) int64 {
	return duration.Milliseconds()
}
