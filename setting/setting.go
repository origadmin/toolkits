package setting

type Setting[S any] func(*S)

// Apply is apply settings
func Apply[T Setting[S], S any](d S, settings ...T) *S {
	val := &d
	for _, setting := range settings {
		setting(val)
	}
	return val
}
