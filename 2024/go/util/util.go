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

// ParseUint8 parses an integer, panics if an error happened.
func ParseUint8(value string) uint8 {
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return uint8(val)
}

// ParseUint64 parses an integer, panics if an error happened.
func ParseUint64(value string) uint64 {
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return uint64(val)
}

// Filter a list, return only elements that return true for the compare method.
func Filter[T any](ts []T, f func(T) bool) []T {
	var result []T
	for _, t := range ts {
		if f(t) {
			result = append(result, t)
		}
	}
	return result
}
