// Generated code by github.com/choonkeat/sumtype-go
package main

// Branch
type branchTreeVariants[T any] struct {
	left  Tree[T]
	right Tree[T]
}

func (s branchTreeVariants[T]) Match(Variants TreeVariants[T]) {
	Variants.Branch(s.left, s.right)
}

func Branch[T any](leftArg Tree[T], rightArg Tree[T]) Tree[T] {
	return branchTreeVariants[T]{leftArg, rightArg}
}

// Leaf
type leafTreeVariants[T any] struct {
	s T
}

func (s leafTreeVariants[T]) Match(Variants TreeVariants[T]) {
	Variants.Leaf(s.s)
}

func Leaf[T any](sArg T) Tree[T] {
	return leafTreeVariants[T]{sArg}
}
