package trees

type BinaryTree struct {
	parent BinaryTraversable
	left   BinaryTraversable
	right  BinaryTraversable
	value  interface{}
}

func NewBinaryTree(value interface{}) *BinaryTree {
	return &BinaryTree{value: value}
}

func (tree *BinaryTree) Parent() BinaryTraversable {
	return tree.parent
}

func (tree *BinaryTree) SetParent(parent BinaryTraversable) {
	tree.parent = parent
}

func (tree *BinaryTree) Left() BinaryTraversable {
	return tree.left
}

func (tree *BinaryTree) SetLeft(left BinaryTraversable) {
	tree.left = left
}

func (tree *BinaryTree) Right() BinaryTraversable {
	return tree.right
}

func (tree *BinaryTree) SetRight(right BinaryTraversable) {
	tree.right = right
}

func (tree *BinaryTree) Value() interface{} {
	return tree.value
}
