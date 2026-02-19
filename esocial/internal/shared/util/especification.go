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

func Not[T any](spec Specification[T]) Specification[T] {
	return func(candidate T) bool {
		return !spec(candidate)
	}
}

func False[T any]() Specification[T] {
	return func(candidate T) bool {
		return false
	}
}

func True[T any]() Specification[T] {
	return func(candidate T) bool {
		return true
	}
}
