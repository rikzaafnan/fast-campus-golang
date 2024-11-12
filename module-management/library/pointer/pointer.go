package pointer

func Of[T any](x T) *T {
	return &x
}

func OfString(x string) *string {
	return &x
}
