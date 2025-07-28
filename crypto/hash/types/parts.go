/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package types

// HashParts represents the parts of a hash
type HashParts struct {
	Algorithm string
	Version   string
	Params    map[string]string
	Hash      []byte
	Salt      []byte
}
