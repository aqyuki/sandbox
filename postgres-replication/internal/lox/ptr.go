package lox

func FromPtr[T any](in *T) T {
	return *in
}
