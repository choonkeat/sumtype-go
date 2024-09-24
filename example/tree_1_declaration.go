package main

// and the variants as a struct
type TreeVariants[T any] struct {
	Branch func(left, right Tree[T])
	Leaf   func(s T)
}
