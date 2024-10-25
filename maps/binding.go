package maps

import (
	"reflect"

	"github.com/goexts/ggb/settings"
)

// DefaultTag is the default tag for binding
const DefaultTag = "json"

// Binder is a binding interface
type Binder interface {
	Bind(typ reflect.Type, tag string, target string)
	Convert(name string) string
}

// Map is a binding map
type Map struct {
	src   string
	binds map[string]Binder // Add a label binds
}

type binds struct {
	values map[string]string
}

// Bind binds a value to a Map.
// The value can be of any type, but the type must be comparable.
// The tag parameter is used to identify the binding in the Map.
// If the type is already bound, the function does nothing.
// The Map will store the binder associated with the type.
func (m *Map) Bind(v interface{}, tag string) {
	// Get the type of the value
	typ := reflect.TypeOf(v)
	// If the value is a pointer, get the element type
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	// If the type is already bound, return
	if _, ok := m.binds[typ.String()]; ok {
		return
	}

	// Create a new binder
	b := NewBinder()
	// Bind the type to the binder
	b.Bind(typ, tag, m.src)
	// Add the binder to the map
	m.binds[typ.String()] = b
}

// Get retrieves the binder associated with the provided value.
// If the value is a pointer, it gets the element type.
// Returns the binder if found in the map, otherwise returns an EmptyBinder.
func (m *Map) Get(v interface{}) Binder {
	// Get the type of the value
	typ := reflect.TypeOf(v)

	// If the value is a pointer, get the element type
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// Check if the type is already bound in the map
	if binder, ok := m.binds[typ.String()]; ok {
		// Return the binder if found
		return binder
	}

	// Return an EmptyBinder if not found
	return EmptyBinder
}

// Bind The label of the binding structure
func (m *binds) Bind(typ reflect.Type, tag string, target string) {
	// Make sure v is a pointer
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// Traverse the fields of the structure
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		name := field.Tag.Get(tag) // Gets the label of the field
		if name == "" {
			name = field.Name
		}
		tag := field.Tag.Get(target)
		if tag == "" {
			tag = field.Name
		}
		if name != tag {
			m.values[name] = tag
		}
	}
}

// Convert returns the converted name if it exists in the binder's values map.
// If the binder is nil or the name is not found, it returns the original name.
func (m *binds) Convert(name string) string {
	// Check if the binder is nil to avoid panicking
	if m == nil {
		return name
	}
	// Check if the name exists in the values map
	if v, ok := m.values[name]; ok {
		// Return the converted name
		return v
	}
	// Return the original name if not found
	return name
}

// NewBinder returns a new instance of the binds struct, which implements the Binder interface.
func NewBinder() Binder {
	// Initialize a new binds struct with an empty values map
	return &binds{
		values: make(map[string]string),
	}
}

// New returns a new instance of the Map struct, which represents a binding map.
// It applies the provided options to the map using the settings.Apply function.
func New(opts ...Option) *Map {
	// Initialize a new Map struct with default values
	return settings.Apply(&Map{
		src:   DefaultTag,
		binds: make(map[string]Binder),
	}, opts)
}

// emptyBinder is a struct that implements the Binder interface, but does nothing.
type emptyBinder struct{}

// Bind is an implementation of the Binder interface that does nothing.
func (e emptyBinder) Bind(typ reflect.Type, tag string, target string) {
	// Do nothing
}

// Convert is an implementation of the Binder interface that returns the original name.
func (e emptyBinder) Convert(name string) string {
	// Return the original name
	return name
}

// EmptyBinder is a variable that holds an instance of the emptyBinder struct, which implements the Binder interface.
var EmptyBinder = Binder(&emptyBinder{})
