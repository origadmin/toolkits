package maps

type Option = func(*Map)

func WithSource(src string) Option {
	return func(m *Map) {
		m.src = src
	}
}
