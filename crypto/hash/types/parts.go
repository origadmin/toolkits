/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package types

// HashParts represents the parts of a hash
type HashParts struct {
	Algorithm Type
	Version   string
	Params    string
	Salt      []byte
	Hash      []byte
}
