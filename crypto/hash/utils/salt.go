/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateSalt generates a random salt with the specified length
func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}
