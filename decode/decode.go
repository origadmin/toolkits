package decode

import (
	"github.com/goexts/generic/maps"
)

// Converter is a reusable object that knows how to convert a collection
// of items (from a map or slice) into a final wrapped structure.
type Converter[I any, W any] struct {
	namer func(name string, item I) (I, bool)
	wrap  func(items []I) W
}

// NewConverter creates a new, reusable converter.
// The namer and wrap functions are defined only once.
func NewConverter[I any, W any](
	namer func(name string, item I) (I, bool),
	wrap func(items []I) W,
) *Converter[I, W] {
	return &Converter[I, W]{
		namer: namer,
		wrap:  wrap,
	}
}

// FromMap executes the conversion from a map.
// It now correctly handles empty maps.
func (c *Converter[I, W]) FromMap(m map[string]I) W {
	// If the map is nil or empty, call the wrapper with nil.
	if len(m) == 0 {
		return c.wrap(nil)
	}

	list := maps.ToSliceWith(m, c.namer)
	return c.wrap(list)
}

// FromSlice executes the conversion from a slice.
// It now correctly handles empty slices, ensuring consistent behavior with FromMap.
func (c *Converter[I, W]) FromSlice(s []I) W {
	// If the slice is nil or empty, call the wrapper with nil.
	if len(s) == 0 {
		return c.wrap(nil)
	}
	return c.wrap(s)
}
