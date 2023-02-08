package alist

func Every[T any](s []T, f func(i int, v T) bool) bool {
	for i, v := range s {
		if !f(i, v) {
			return false
		}
	}
	return true
}

func Some[T any](s []T, f func(i int, v T) bool) bool {
	for i, v := range s {
		if f(i, v) {
			return true
		}
	}
	return false
}

func Filter[T any](s []T, f func(i int, v T) bool) []T {
	if s == nil {
		return nil
	}
	a := make([]T, 0, len(s))
	for i, v := range s {
		if f(i, v) {
			a = append(a, v)
		}
	}
	return a[:len(a):len(a)]
}

func Reduce[T, U any](s []T, f func(i int, v T, prev U) U, initial U) U {
	state := initial
	for i, v := range s {
		state = f(i, v, state)
	}
	return state
}

func ReduceSimple[T any](s []T, f func(i int, v T, prev T) T) T {
	if len(s) == 0 {
		panic("empty slice")
	}
	return Reduce(s[1:], f, s[0])
}
