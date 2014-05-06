package trees

import (
	"github.com/modocache/cargo/comparators"
	"reflect"
)

type BinarySearchTree struct {
	Left  *BinarySearchTree
	Right *BinarySearchTree
	Value interface{}
	less  comparators.Less
}

func NewBinarySearchTree(value interface{}, less comparators.Less) *BinarySearchTree {
	return &BinarySearchTree{Value: value, less: less}
}

func (tree *BinarySearchTree) Insert(value interface{}) {
	if tree.less(value, tree.Value) {
		tree.insertLeft(value)
	} else {
		tree.insertRight(value)
	}
}

func (tree *BinarySearchTree) insertLeft(value interface{}) {
	if tree.Left == nil {
		tree.Left = NewBinarySearchTree(value, tree.less)
	} else {
		tree.Left.Insert(value)
	}
}

func (tree *BinarySearchTree) insertRight(value interface{}) {
	if tree.Right == nil {
		tree.Right = NewBinarySearchTree(value, tree.less)
	} else {
		tree.Right.Insert(value)
	}
}

func (tree *BinarySearchTree) Find(value interface{}) *BinarySearchTree {
	if reflect.DeepEqual(value, tree.Value) {
		return tree
	} else if tree.less(value, tree.Value) {
		return tree.findLeft(value)
	} else {
		return tree.findRight(value)
	}
}

func (tree *BinarySearchTree) findLeft(value interface{}) *BinarySearchTree {
	if tree.Left == nil {
		return nil
	} else {
		return tree.Left.Find(value)
	}
}

func (tree *BinarySearchTree) findRight(value interface{}) *BinarySearchTree {
	if tree.Right == nil {
		return nil
	} else {
		return tree.Right.Find(value)
	}
}
