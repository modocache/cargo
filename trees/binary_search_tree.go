package trees

import (
	"github.com/modocache/cargo/comparators"
)

type BinarySearchTree struct {
	*BinaryTree
	less comparators.Less
}

func NewBinarySearchTree(value interface{}, less comparators.Less) *BinarySearchTree {
	return &BinarySearchTree{NewBinaryTree(value), less}
}

func (tree *BinarySearchTree) Less() comparators.Less {
	return tree.less
}

func (tree *BinarySearchTree) Insert(value interface{}) BinarySearchable {
	constructor := func(parent BinarySearchable, value interface{}) BinarySearchable {
		binaryTree := NewBinaryTree(value)
		binaryTree.SetParent(parent)
		return &BinarySearchTree{binaryTree, tree.less}
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
