package trees

import (
	"github.com/modocache/cargo/comparators"
)

type RedBlackTreeColor int

const (
	Red   RedBlackTreeColor = 1
	Black RedBlackTreeColor = 2
)

func setColor(tree *RedBlackTree, color RedBlackTreeColor) {
	if tree != nil {
		tree.color = color
	}
}

type RedBlackTree struct {
	*BinarySearchTree
	color RedBlackTreeColor
}

func NewRedBlackTree(value interface{}, less comparators.Less) *RedBlackTree {
	return &RedBlackTree{&BinarySearchTree{value: value, less: less}, Black}
}

func (tree *RedBlackTree) Insert(value interface{}) BinarySearchable {
	constructor := func(parent BinarySearchable, value interface{}) BinarySearchable {
		return &RedBlackTree{&BinarySearchTree{value: value, less: tree.less, parent: parent}, Red}
	}

	inserted := insert(tree, value, constructor).(*RedBlackTree)
	inserted.balance()
	return inserted
}

func (tree *RedBlackTree) balance() {
	for tree.parent() != nil && tree.parent().color == Red {
		if isParentLeftChild(tree) {
			tree = tree.balanceLeftChild()
		} else {
			tree = tree.balanceRightChild()
		}
	}

	Root(tree).(*RedBlackTree).color = Black
}

func (tree *RedBlackTree) balanceLeftChild() *RedBlackTree {
	if rightUncle := tree.rightUncle(); rightUncle != nil && rightUncle.color == Red {
		// If tree has a right uncle, then the subtree with
		// tree.grandParent() at its root is balanced. Reset the colors
		// to maintain red-black tree conditions (i.e.: red nodes may not
		// have red children), then balance the grandParent.
		setColor(rightUncle, Black)
		setColor(tree.parent(), Black)
		setColor(tree.grandParent(), Red)
		tree = tree.grandParent()
	} else if isRightChild(tree) {
		tree = tree.parent()
		rotateLeft(tree)
		setColor(tree.parent(), Black)
		setColor(tree.grandParent(), Red)
		tree.rotateGrandParentRight()
	}
	return tree
}

func (tree *RedBlackTree) balanceRightChild() *RedBlackTree {
	// This is the right-sided version of .balanceLeftChild().
	// Detailed comments may be found there.
	if leftUncle := tree.leftUncle(); leftUncle != nil && leftUncle.color == Red {
		setColor(leftUncle, Black)
		setColor(tree.parent(), Black)
		setColor(tree.grandParent(), Red)
		tree = tree.grandParent()
	} else if isLeftChild(tree) {
		tree = tree.parent()
		rotateRight(tree)
		setColor(tree.parent(), Black)
		setColor(tree.grandParent(), Red)
		tree.rotateGrandParentLeft()
	}
	return tree
}

func (tree *RedBlackTree) parent() *RedBlackTree {
	if tree.Parent() == nil {
		return nil
	} else {
		return tree.Parent().(*RedBlackTree)
	}
}

func (tree *RedBlackTree) grandParent() *RedBlackTree {
	if tree.parent() == nil {
		return nil
	} else {
		return tree.parent().parent()
	}
}

func (tree *RedBlackTree) rightUncle() *RedBlackTree {
	if uncle := rightUncle(tree); uncle == nil {
		return nil
	} else {
		return uncle.(*RedBlackTree)
	}
}

func (tree *RedBlackTree) leftUncle() *RedBlackTree {
	if uncle := leftUncle(tree); uncle == nil {
		return nil
	} else {
		return uncle.(*RedBlackTree)
	}
}

func (tree *RedBlackTree) rotateGrandParentRight() {
	if tree.grandParent() != nil {
		rotateRight(tree.grandParent())
	}
}

func (tree *RedBlackTree) rotateGrandParentLeft() {
	if tree.grandParent() != nil {
		rotateLeft(tree.grandParent())
	}
}
