package generic

func Zero[T any]() T {
	var zero T
	return zero
}
