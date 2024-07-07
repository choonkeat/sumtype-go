// Generated code by github.com/choonkeat/sumtype-go
package main

// Branch
type branchTreeVariants[T any] struct {
	left  Tree[T]
	right Tree[T]
}

func (s branchTreeVariants[T]) Match(variants TreeVariants[T]) {
	variants.Branch(s.left, s.right)
}

// Branch is a constructor function for Tree; see TreeVariants for all constructor functions of Tree
func Branch[T any](leftArg Tree[T], rightArg Tree[T]) Tree[T] {
	return branchTreeVariants[T]{leftArg, rightArg}
}

// Leaf
type leafTreeVariants[T any] struct {
	s T
}

func (s leafTreeVariants[T]) Match(variants TreeVariants[T]) {
	variants.Leaf(s.s)
}

// Leaf is a constructor function for Tree; see TreeVariants for all constructor functions of Tree
func Leaf[T any](sArg T) Tree[T] {
	return leafTreeVariants[T]{sArg}
}

// TreeVariantsMap is parameter type of TreeMap function,
// like TreeVariants is parameter type of tree.Match method,
// but with methods that returns a value of generic type
type TreeVariantsMap[T, A any] struct {
	Branch func(leftArg Tree[T], rightArg Tree[T]) A // when Tree value pattern matches to Branch, return different value
	Leaf   func(sArg T) A                            // when Tree value pattern matches to Leaf, return different value
}

// TreeMap is like tree.Match method except it returns a value of generic type
// thus can transform a Tree value into anything else
func TreeMap[T, A any](value Tree[T], variants TreeVariantsMap[T, A]) A {
	var result A
	value.Match(TreeVariants[T]{
		Branch: func(leftArg Tree[T], rightArg Tree[T]) {
			result = variants.Branch(leftArg, rightArg)
		},
		Leaf: func(sArg T) {
			result = variants.Leaf(sArg)
		},
	})
	return result
}
