package main

type Tree[T interface{}] interface {
	Match(s TreeVariants[T])
}

// and the variants as a struct
type TreeVariants[T interface{}] struct {
	Branch func(left, right Tree[T])
	Leaf   func(s T)
}
