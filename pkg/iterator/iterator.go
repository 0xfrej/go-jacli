package iterator

type Iterator[E any] interface {
	HasNext() bool
	Next() (elem E, ok bool)
	Peek() (elem E, ok bool)
}
