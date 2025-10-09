package utils

func From[T any](v T) *T {
	return &v
}

func Deref[T any](ptr *T) T {
	if ptr != nil {
		return *ptr
	}
	var d T
	return d
}
