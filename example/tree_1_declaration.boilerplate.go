// Generated code by github.com/choonkeat/sumtype-go
package main

import (
	"encoding/json"
	"fmt"
)

// Branch
type branchTreeVariants[T any] struct {
	Left  Tree[T]
	Right Tree[T]
}

func (s branchTreeVariants[T]) Match(variants TreeVariants[T]) {
	variants.Branch(s.Left, s.Right)
}

// Branch is a constructor function for Tree; see TreeVariants for all constructor functions of Tree
func Branch[T any](leftArg Tree[T], rightArg Tree[T]) Tree[T] {
	return Tree[T]{branchTreeVariants[T]{leftArg, rightArg}}
}

// Leaf
type leafTreeVariants[T any] struct {
	S T
}

func (s leafTreeVariants[T]) Match(variants TreeVariants[T]) {
	variants.Leaf(s.S)
}

// Leaf is a constructor function for Tree; see TreeVariants for all constructor functions of Tree
func Leaf[T any](sArg T) Tree[T] {
	return Tree[T]{leafTreeVariants[T]{sArg}}
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

// Tree = Branch | Leaf
type Tree[T any] struct {
	tree tree[T]
}

// tree is the interface for TreeVariants
type tree[T any] interface {
	Match(variants TreeVariants[T])
}

func (s Tree[T]) Match(variants TreeVariants[T]) {
	s.tree.Match(variants)
}
func (s Tree[T]) MarshalJSON() (data []byte, err error) {
	s.tree.Match(TreeVariants[T]{
		Branch: func(leftArg Tree[T], rightArg Tree[T]) {
			data, err = json.Marshal([]any{
				"Branch",
				branchTreeVariants[T]{
					Left:  leftArg,
					Right: rightArg,
				}})
		},
		Leaf: func(sArg T) {
			data, err = json.Marshal([]any{
				"Leaf",
				leafTreeVariants[T]{
					S: sArg,
				}})
		},
	})
	return data, err
}
func (s *Tree[T]) UnmarshalJSON(data []byte) error {
	// The expected format is ["TypeName", { ... data... }]
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("expected an array with type and data, got error: %w", err)
	}
	if len(raw) != 2 {
		return fmt.Errorf("expected array of two elements [type, data], got %d elements", len(raw))
	}
	// Unmarshal the first element to get the type
	var typeName string
	if err := json.Unmarshal(raw[0], &typeName); err != nil {
		return fmt.Errorf("failed to unmarshal type name: %w", err)
	}
	switch typeName {
	case "Branch":
		var temp branchTreeVariants[T]
		if err := json.Unmarshal(raw[1], &temp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		s.tree = branchTreeVariants[T]{
			Left:  temp.Left,
			Right: temp.Right,
		}
	case "Leaf":
		var temp leafTreeVariants[T]
		if err := json.Unmarshal(raw[1], &temp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		s.tree = leafTreeVariants[T]{
			S: temp.S,
		}
	default:
		return fmt.Errorf("unknown type %q", typeName)
	}
	return nil
}
