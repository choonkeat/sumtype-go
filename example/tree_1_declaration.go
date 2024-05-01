package main

type Tree[T any] interface {
	Match(s TreeVariants[T])
}

// and the variants as a struct
type TreeVariants[T any] struct {
	Branch func(left, right Tree[T])
	Leaf   func(s T)
}
