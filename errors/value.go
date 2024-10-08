package errors

type Valuer[T any] interface {
	Value() T
	Error() string
	Unwrap() error
	Is(err error) bool
}

type wrapped[T any] struct {
	v T
	e error
}

func Valued[T any](v T, err error) Valuer[T] {
	return wrapped[T]{v, err}
}

func (obj wrapped[T]) Value() T {
	return obj.v
}

func (obj wrapped[T]) Unwrap() error {
	return obj.e
}

func (obj wrapped[T]) Error() string {
	if obj.e == nil {
		return ""
	}
	return obj.e.Error()
}

func (obj wrapped[T]) Is(err error) bool {
	if obj.e == nil {
		return err == nil
	}

	return Is(obj.e, err)
}
