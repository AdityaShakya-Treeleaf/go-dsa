package utils

func TernaryOperation[T any](condition bool, first, second T) T {
	if condition {
		return first
	} else {
		return second
	}
}
