package trees

import (
	"github.com/modocache/cargo/comparators"
)

type BinarySearchTree struct {
	parent BinarySearchable
	left   BinarySearchable
	right  BinarySearchable
	value  interface{}
	less   comparators.Less
}

func NewBinarySearchTree(value interface{}, less comparators.Less) *BinarySearchTree {
	return &BinarySearchTree{value: value, less: less}
}

func (tree *BinarySearchTree) Parent() BinarySearchable {
	return tree.parent
}

func (tree *BinarySearchTree) SetParent(parent BinarySearchable) {
	tree.parent = parent
}

func (tree *BinarySearchTree) Left() BinarySearchable {
	return tree.left
}

func (tree *BinarySearchTree) SetLeft(left BinarySearchable) {
	tree.left = left
}

func (tree *BinarySearchTree) Right() BinarySearchable {
	return tree.right
}

func (tree *BinarySearchTree) SetRight(right BinarySearchable) {
	tree.right = right
}

func (tree *BinarySearchTree) Value() interface{} {
	return tree.value
}

func (tree *BinarySearchTree) Less() comparators.Less {
	return tree.less
}

func (tree *BinarySearchTree) Insert(value interface{}) BinarySearchable {
	constructor := func(parent BinarySearchable, value interface{}) BinarySearchable {
		return &BinarySearchTree{value: value, less: tree.less, parent: parent}
	}
	return insert(tree, value, constructor).(*BinarySearchTree)
}

func (tree *BinarySearchTree) InsertAll(values ...interface{}) {
	for _, value := range values {
		tree.Insert(value)
	}
}

func (tree *BinarySearchTree) Find(value interface{}) BinarySearchable {
	return find(tree, value)
}
