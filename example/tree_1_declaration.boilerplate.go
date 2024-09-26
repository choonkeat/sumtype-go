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

func (branchInstance branchTreeVariants[T]) Match(branchVariants TreeVariants[T]) {
	branchVariants.Branch(branchInstance.Left, branchInstance.Right)
}

// Branch is a constructor function for Tree; see TreeVariants for all constructor functions of Tree
func Branch[T any](leftArg Tree[T], rightArg Tree[T]) Tree[T] {
	return Tree[T]{branchTreeVariants[T]{leftArg, rightArg}}
}

// Leaf
type leafTreeVariants[T any] struct {
	S T
}

func (leafInstance leafTreeVariants[T]) Match(leafVariants TreeVariants[T]) {
	leafVariants.Leaf(leafInstance.S)
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
func TreeMap[T, A any](treeValue Tree[T], treeVariants TreeVariantsMap[T, A]) A {
	var treeTemp A
	treeValue.Match(TreeVariants[T]{
		Branch: func(leftArg Tree[T], rightArg Tree[T]) {
			treeTemp = treeVariants.Branch(leftArg, rightArg)
		},
		Leaf: func(sArg T) {
			treeTemp = treeVariants.Leaf(sArg)
		},
	})
	return treeTemp
}

// Tree = Branch | Leaf
type Tree[T any] struct {
	tree tree[T]
}

// tree is the interface for TreeVariants
type tree[T any] interface {
	Match(variants TreeVariants[T])
}

func (treeInstance Tree[T]) Match(treeVariants TreeVariants[T]) {
	treeInstance.tree.Match(treeVariants)
}
func (treeInstance Tree[T]) MarshalJSON() (treeData []byte, treeErr error) {
	treeInstance.tree.Match(TreeVariants[T]{
		Branch: func(leftArg Tree[T], rightArg Tree[T]) {
			treeData, treeErr = json.Marshal([]any{
				"Branch",
				branchTreeVariants[T]{
					Left:  leftArg,
					Right: rightArg,
				}})
		},
		Leaf: func(sArg T) {
			treeData, treeErr = json.Marshal([]any{
				"Leaf",
				leafTreeVariants[T]{
					S: sArg,
				}})
		},
	})
	return treeData, treeErr
}
func (treeInstance *Tree[T]) UnmarshalJSON(treeData []byte) error {
	// The expected format is ["TypeName", { ... data... }]
	var treeRaw []json.RawMessage
	if err := json.Unmarshal(treeData, &treeRaw); err != nil {
		return fmt.Errorf("expected an array with type and data, got error: %w", err)
	}
	if len(treeRaw) != 2 {
		return fmt.Errorf("expected array of two elements [type, data], got %d elements", len(treeRaw))
	}
	// Unmarshal the first element to get the type
	var treeVariantName string
	if err := json.Unmarshal(treeRaw[0], &treeVariantName); err != nil {
		return fmt.Errorf("failed to unmarshal type name: %w", err)
	}
	switch treeVariantName {
	case "Branch":
		var treeTemp branchTreeVariants[T]
		if err := json.Unmarshal(treeRaw[1], &treeTemp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		treeInstance.tree = branchTreeVariants[T]{
			Left:  treeTemp.Left,
			Right: treeTemp.Right,
		}
	case "Leaf":
		var treeTemp leafTreeVariants[T]
		if err := json.Unmarshal(treeRaw[1], &treeTemp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		treeInstance.tree = leafTreeVariants[T]{
			S: treeTemp.S,
		}
	default:
		return fmt.Errorf("unknown type %q", treeVariantName)
	}
	return nil
}
