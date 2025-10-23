/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package identifier

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// --- Test Helpers ---

// mockProvider is a helper for testing registration.
type mockProvider struct {
	name       string
	isStringer bool
	isNumberer bool
}

func (m *mockProvider) Name() string { return m.name }
func (m *mockProvider) Size() int    { return 1 }
func (m *mockProvider) AsString() Generator[string] {
	if m.isStringer {
		return &mockStringGenerator{m}
	}
	return nil
}
func (m *mockProvider) AsNumber() Generator[int64] {
	if m.isNumberer {
		return &mockNumberGenerator{m}
	}
	return nil
}

// mockStringGenerator is a mock string generator.
type mockStringGenerator struct {
	*mockProvider
}

func (m *mockStringGenerator) Generate() string        { return m.name + "_string" }
func (m *mockStringGenerator) Validate(id string) bool { return id == m.name+"_string" }

// mockNumberGenerator is a mock number generator.
type mockNumberGenerator struct {
	*mockProvider
}

func (m *mockNumberGenerator) Generate() int64        { return 42 }
func (m *mockNumberGenerator) Validate(id int64) bool { return id == 42 }

// resetRegistry is a crucial helper to ensure test isolation.
func resetRegistry() {
	globalRegistry = &registry{
		providers: map[string]Provider{
			defaultString: builtinString,
			defaultNumber: builtinNumber,
		},
		defaultStringName: defaultString,
		defaultNumberName: defaultNumber,
	}
}

// --- Tests ---

func TestOutOfTheBoxDefaults(t *testing.T) {
	resetRegistry()

	// Test default string generation (should be built-in UUID)
	strID := GenerateString()
	assert.NotEmpty(t, strID)
	assert.True(t, Get[string](defaultString).Validate(strID), "default string ID should be a valid UUID")

	// Test default number generation (should be built-in Snowflake)
	numID := GenerateNumber()
	assert.NotZero(t, numID)
	assert.True(t, Get[int64](defaultNumber).Validate(numID), "default number ID should be a valid Snowflake ID")
}

func TestGet(t *testing.T) {
	resetRegistry()

	// Test getting pre-registered providers
	assert.NotNil(t, Get[string]("uuid"))
	assert.NotNil(t, Get[int64]("snowflake"))

	// Test getting a non-existent provider
	assert.Nil(t, Get[string]("non-existent"))

	// Test getting with wrong type
	assert.Nil(t, Get[int64]("uuid"), "uuid should not have a number generator")
	assert.Nil(t, Get[string]("snowflake"), "snowflake should not have a string generator")
}

func TestRegisterAndOverride(t *testing.T) {
	resetRegistry()

	// 1. Register a new provider
	newStringer := &mockProvider{name: "cuid2", isStringer: true}
	Register(newStringer)
	assert.NotNil(t, Get[string]("cuid2"))
	assert.Equal(t, "cuid2_string", Get[string]("cuid2").Generate())

	// 2. Override a built-in provider
	newUUID := &mockProvider{name: "uuid", isStringer: true}
	Register(newUUID)
	assert.Equal(t, "uuid_string", GenerateString(), "GenerateString should now use the new override")

	// 3. Test panic on nil registration
	assert.Panics(t, func() {
		Register(nil)
	})
}

func TestSetDefaults(t *testing.T) {
	resetRegistry()

	// 1. Register new providers
	cuid2 := &mockProvider{name: "cuid2", isStringer: true}
	sonyflake := &mockProvider{name: "sonyflake", isNumberer: true}
	Register(cuid2)
	Register(sonyflake)

	// 2. Set new defaults
	SetDefaultString("cuid2")
	SetDefaultNumber("sonyflake")

	// 3. Verify new defaults are used
	assert.Equal(t, "cuid2_string", GenerateString())
	assert.Equal(t, int64(42), GenerateNumber())

	// 4. Test panic on setting unregistered default
	assert.Panics(t, func() {
		SetDefaultString("non-existent")
	})
}

func TestConcurrency(t *testing.T) {
	resetRegistry()

	var wg sync.WaitGroup
	const numRoutines = 100

	// Concurrently generate IDs
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = GenerateString()
			_ = GenerateNumber()
		}()
	}

	// Concurrently register and set defaults
	for i := 0; i < numRoutines/10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// Register new providers with unique names
			name := fmt.Sprintf("mock-%d", i)
			Register(&mockProvider{name: name, isStringer: true})
			// Occasionally set a new default
			if i%5 == 0 {
				SetDefaultString(name)
			}
		}(i)
	}

	wg.Wait()
	// The test passes if it completes without panicking due to race conditions.
}
