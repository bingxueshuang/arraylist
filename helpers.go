package alist

import "golang.org/x/exp/constraints"

func Equal[T comparable](x, y T) bool {
	return x == y
}

func Less[T constraints.Ordered](x, y T) bool {
	return x < y
}

func IsSorted[T any](s []T, less func(x, y T) bool) bool {
	for i := len(s) - 1; i > 0; i-- {
		if less(s[i], s[i-1]) {
			return false
		}
	}
	return true
}
