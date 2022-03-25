package masochism

func Foreach[T any](f func(elem T)) func(iterable []T) {
	return func(iterable []T) {
		for i := 0; i < len(iterable); i++ {
			f(iterable[i])
		}
	}
}
