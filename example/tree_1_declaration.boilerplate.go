// Generated code by github.com/choonkeat/sumtype-go
package main

// Branch
type branchTreeVariants[T interface{}] struct {
	left  Tree[T]
	right Tree[T]
}

func (s branchTreeVariants[T]) Match(Variants TreeVariants[T]) {
	Variants.Branch(s.left, s.right)
}

func Branch[T interface{}](leftArg Tree[T], rightArg Tree[T]) Tree[T] {
	return branchTreeVariants[T]{leftArg, rightArg}
}

// Leaf
type leafTreeVariants[T interface{}] struct {
	s T
}

func (s leafTreeVariants[T]) Match(Variants TreeVariants[T]) {
	Variants.Leaf(s.s)
}

func Leaf[T interface{}](sArg T) Tree[T] {
	return leafTreeVariants[T]{sArg}
}
