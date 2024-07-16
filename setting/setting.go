package setting

// Apply is apply settings
func Apply[T func(*S), S any](d S, settings ...T) *S {
	val := &d
	for _, setting := range settings {
		setting(val)
	}
	return val
}
