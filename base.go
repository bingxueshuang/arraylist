package alist

type List[T any] []T

func (a List[T]) Clone() List[T] {
	if a == nil {
		return nil
	}
	var s List[T] = make([]T, len(a))
	copy(s, a)
	return s
}

func (a List[T]) Contains(item T, eq func(x, y T) bool) bool {
	return a.Index(item, eq) >= 0
}

func (a List[T]) Index(item T, eq func(x, y T) bool) int {
	for i, v := range a {
		if eq(v, item) {
			return i
		}
	}
	return -1
}

func (a List[T]) Append(items ...T) List[T] {
	s := a.growTo(len(a) + len(items))
	return append(s, items...)
}

func (a List[T]) Splice(i, j int, items ...T) List[T] {
	_ = a[i:j]
	m := len(items)
	n := len(a[:i]) + m + len(a[j:])
	s := a.growTo(n)
	if n > len(s) {
		s = s[:n]
	}
	copy(s[i+m:], s[j:])
	copy(s[i:j], items)
	s = s[:n]
	return s.shrink()
}

func (a List[T]) growTo(n int) List[T] {
	c := cap(a)
	if n <= c {
		return a
	}
	for n > c {
		c = c * 2
	}
	s := make([]T, len(a), c)
	copy(s, a)
	return s
}

func (a List[T]) shrink() List[T] {
	n := len(a)
	c := cap(a)
	if n <= c/4 {
		return a[: n : 2*n]
	}
	return a
}
