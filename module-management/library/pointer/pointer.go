package pointer

func Of[T any](x T) *T {
	return &x
}
