/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package xid_test // Use black-box testing

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	// Blank import to trigger the xid provider registration
	_ "github.com/origadmin/toolkits/identifier/xid"
)

// TestXidProvider ensures the xid provider is registered correctly.
func TestXidProvider(t *testing.T) {
	// String generator should be available.
	strGenerator := identifier.New[string]("xid")
	assert.NotNil(t, strGenerator, "Expected to get a non-nil string generator for 'xid'")

	// Number generator should NOT be available.
	numGenerator := identifier.New[int64]("xid")
	assert.Nil(t, numGenerator, "Expected to get a nil number generator for 'xid' as it is not supported")
}

// TestGenerateAndValidate tests the generation and validation of an XID.
func TestGenerateAndValidate(t *testing.T) {
	generator := identifier.New[string]("xid")
	if !assert.NotNil(t, generator, "Generator should not be nil") {
		t.FailNow()
	}

	// 1. Generate a new ID
	id := generator.Generate()
	assert.NotEmpty(t, id, "Generated ID should not be empty")
	assert.Len(t, id, 20, "A standard XID should have a length of 20 characters")

	// 2. Validate the new ID
	isValid := generator.Validate(id)
	assert.True(t, isValid, "A freshly generated ID should be valid")

	// 3. Validate a known invalid ID
	isInvalid := generator.Validate("not-a-valid-xid")
	assert.False(t, isInvalid, "An invalid string should not be considered a valid XID")
}

// TestGeneratorProperties checks the metadata of the xid generator.
func TestGeneratorProperties(t *testing.T) {
	generator := identifier.New[string]("xid")
	if !assert.NotNil(t, generator, "Generator should not be nil") {
		t.FailNow()
	}

	assert.Equal(t, "xid", generator.Name(), "Generator name should be 'xid'")
	assert.Equal(t, 96, generator.Size(), "Generator size should be 96 bits")
}
