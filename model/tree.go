package model

import (
	"reflect"
	"sort"
)

// Tree is a binary tree
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree sending all values from the tree to the channel ch
func (tree *Tree) Walk(ch chan int) {
	tree.walk(ch)
	close(ch)
}

func (tree *Tree) walk(ch chan int) {
	ch <- tree.Value
	if tree.Left != nil {
		tree.Left.walk(ch)
	}
	if tree.Right != nil {
		tree.Right.walk(ch)
	}
}

// Same determines whether the trees contain the same values
func (tree *Tree) Same(tree2 *Tree) bool {
	slice1 := tree.toSlice()
	slice2 := tree2.toSlice()
	sort.Slice(slice1, func(i, j int) bool { return slice1[i] < slice1[j] })
	sort.Slice(slice2, func(i, j int) bool { return slice2[i] < slice2[j] })
	return reflect.DeepEqual(slice1, slice2)
}

func (tree *Tree) toSlice() []int {
	var slice []int
	ch := make(chan int)
	go tree.Walk(ch)
	for i := range ch {
		slice = append(slice, i)
	}
	return slice
}

// NewTree makes a tree
// TODO: have it construct a binary tree from elems
func NewTree(n int) *Tree {
	return &Tree{
		Left: &Tree{
			Left:  nil,
			Value: 4 * n,
			Right: nil,
		},
		Value: 7 * n,
		Right: &Tree{
			Left: &Tree{
				Left:  nil,
				Value: 8 * n,
				Right: nil,
			},
			Value: 9 * n,
			Right: &Tree{
				Left:  nil,
				Value: 10 * n,
				Right: nil,
			},
		},
	}
}
