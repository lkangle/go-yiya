package utils

func Map[T any, R any](arr []T, f func(T) R) []R {
	result := make([]R, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}
