// Generated code by github.com/choonkeat/sumtype-go
package main

// Branch
type branchTreeScenarios[T interface{}] struct {
	left  Tree[T]
	right Tree[T]
}

func (s branchTreeScenarios[T]) Switch(scenarios TreeScenarios[T]) {
	scenarios.Branch(s.left, s.right)
}

func Branch[T interface{}](leftArg Tree[T], rightArg Tree[T]) Tree[T] {
	return branchTreeScenarios[T]{leftArg, rightArg}
}

// Leaf
type leafTreeScenarios[T interface{}] struct {
	s T
}

func (s leafTreeScenarios[T]) Switch(scenarios TreeScenarios[T]) {
	scenarios.Leaf(s.s)
}

func Leaf[T interface{}](sArg T) Tree[T] {
	return leafTreeScenarios[T]{sArg}
}
