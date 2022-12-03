package slice_functions

// Map maps slice to another slice of different types
func Map[T any, V any, F func(T) V](slice []T, f F) []V {
	res := make([]V, 0)
	for _, t := range slice {
		res = append(res, f(t))
	}
	return res
}

// FlatMap maps slice to another slice of different types flattening the result of f
func FlatMap[T any, V any, F func(T) []V](slice []T, f F) []V {
	res := make([]V, 0)
	for _, t := range slice {
		res = append(res, f(t)...)
	}
	return res
}

// FlatMapTry maps slice to another slice of different types flattening the result of f with possible errors
func FlatMapTry[T any, V any, F func(T) ([]V, error)](slice []T, f F) ([]V, error) {
	res := make([]V, 0)
	for _, t := range slice {
		fs, err := f(t)
		if err != nil {
			return nil, err
		}
		res = append(res, fs...)
	}
	return res, nil
}

// Filter filters slice by condition
func Filter[T any, F func(T) bool](slice []T, f F) []T {
	res := make([]T, 0)
	for _, t := range slice {
		if f(t) {
			res = append(res, t)
		}
	}
	return res
}

// Count counts number of elements satisfying the condition
func Count[T any, F func(T) bool](slice []T, f F) int {
	res := 0
	for _, t := range slice {
		if f(t) {
			res++
		}
	}
	return res
}

// ForEach applies function to each element of slice
func ForEach[T any, F func(T)](slice []T, f F) {
	for _, t := range slice {
		f(t)
	}
}

// ForAll returns true if f condition is satisfied for all elements
func ForAll[T any](ts []T, f func(T) bool) bool {
	for _, t := range ts {
		if !f(t) {
			return false
		}
	}
	return true
}

// Exists returns true if f condition is satisfied for one of the elements
func Exists[T any](ts []T, f func(T) bool) bool {
	for _, t := range ts {
		if f(t) {
			return true
		}
	}
	return false
}

// GroupBy creates a map from slice based on key, values from pairs of function result
func GroupBy[T any, K comparable, V any](ts []T, f func(T) (K, V)) map[K][]V {
	res := make(map[K][]V)
	for _, t := range ts {
		k, v := f(t)
		res[k] = append(res[k], v)
	}
	return res
}

// MapBy creates a map from slice based on key, values from pairs of function result
// assuming each key only has one value
func MapBy[T any, K comparable, V any](ts []T, f func(T) (K, V)) map[K]V {
	res := make(map[K]V)
	for _, t := range ts {
		k, v := f(t)
		res[k] = v
	}
	return res
}

// New creates a new slice typed by argument types
func New[T any](ts ...T) []T {
	return ts
}
