/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package rand

// DefaultSaltSize is the default length of the salt string.
const DefaultSaltSize = 16

// GenerateRandom generates a random salt string of specified length.
//
// Parameters:
//
//	length int - the desired length of the salt string
//
// Return value:
//
//	string - the generated random salt string
func GenerateRandom(length int) string {
	return string(All.RandBytes(length))
}

// GenerateSalt generates a random salt string of specified length.
//
// Return value:
//
//	string - the generated random salt string
func GenerateSalt() string {
	return string(All.RandBytes(DefaultSaltSize))
}
