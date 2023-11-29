package main

type Tree[T interface{}] interface {
	Switch(s TreeScenarios[T])
}

// and the variants as a struct
type TreeScenarios[T interface{}] struct {
	Branch func(left, right Tree[T])
	Leaf   func(s T)
}
