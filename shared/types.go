package types

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

type Tuple_2[T1 any, T2 any] struct {
	One T1
	Two T2
}
