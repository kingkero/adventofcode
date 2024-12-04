package util

import (
	"strconv"
)

// Map returns a list of elements, where the new element is the result of
// running the passed function.
func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

// ParseInt parses an integer, panics if an error happened.
func ParseInt(value string) int {
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(val)
}
