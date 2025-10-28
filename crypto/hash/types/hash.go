// Package types implements the functions, types, and interfaces for the module.
package types

import (
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
)

// Hash is a helper function that might need to be refactored
// depending on how stdhash.ParseHash is updated to handle the new Spec struct.
// For now, it assumes subAlg is a simple string.
func Hash(alg string) (stdhash.Hash, error) {
	h, err := stdhash.ParseHash(alg)
	if err != nil {
		return 0, fmt.Errorf("unsupported hash type: %s", alg)
	}
	return h, nil
}
