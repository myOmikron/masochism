package masochism

func Foreach[T any](f func(elem T)) func(iterable []T) {
	return func(iterable []T) {
		for i := 0; i < len(iterable); i++ {
			f(iterable[i])
		}
	}
}

func Map[T []E, E any](f func(elem E) E) func(iterable T) T {
	return func(iterable T) T {
		r := make([]E, 0)
		for _, elem := range iterable {
			r = append(r, f(elem))
		}
		return r
	}
}

func Filter[T []E, E any](filter func(elem E) bool) func(iterable T) T {
	return func(iterable T) T {
		r := make([]E, 0)
		for _, elem := range iterable {
			if filter(elem) {
				r = append(r, elem)
			}
		}
		return r
	}
}
