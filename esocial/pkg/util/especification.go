package util

// Specification como função
type Specification[T any] func(T) bool

// Combinadores
func And[T any](specs ...Specification[T]) Specification[T] {
	return func(candidate T) bool {
		for _, spec := range specs {
			if !spec(candidate) {
				return false
			}
		}
		return true
	}
}

func Or[T any](specs ...Specification[T]) Specification[T] {
	return func(candidate T) bool {
		for _, spec := range specs {
			if spec(candidate) {
				return true
			}
		}
		return false
	}
}
