package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	value int
}

type BinaryTree struct {
	root *Node
}

func (t *BinaryTree) Insert(value int) {
	if t.root == nil {
		t.root = &Node{value: value}
		return
	}
	t.insertNode(t.root, &Node{value: value})
}

func (t *BinaryTree) insertNode(parent, newNode *Node) {
	if newNode.value < parent.value {
		if parent.left == nil {
			parent.left = newNode
		} else {
			t.insertNode(parent.left, newNode)
		}
		return
	}
	if parent.right == nil {
		parent.right = newNode
	} else {
		t.insertNode(parent.right, newNode)
	}
}

// func (t *BinaryTree) Search(value int) bool              {}
func (t *BinaryTree) TraverseInOrder(fn func(value int)) {
	t.recursiveInOrder(t.root, fn)
}

func (t *BinaryTree) recursiveInOrder(node *Node, fn func(value int)) {
	if node == nil {
		return
	}
	t.recursiveInOrder(node.left, fn)
	fn(node.value)
	t.recursiveInOrder(node.right, fn)
}

func (t *BinaryTree) TraversePreOrder(fn func(value int)) {
	t.recursivePreOrder(t.root, fn)
}

func (t *BinaryTree) recursivePreOrder(node *Node, fn func(value int)) {
	if node == nil {
		return
	}
	fn(node.value)
	t.recursiveInOrder(node.left, fn)
	t.recursiveInOrder(node.right, fn)
}

func (t *BinaryTree) TraversePostOrder(fn func(value int)) {
	t.recursivePostOrder(t.root, fn)
}

func (t *BinaryTree) recursivePostOrder(node *Node, fn func(value int)) {
	if node == nil {
		return
	}
	t.recursiveInOrder(node.left, fn)
	t.recursiveInOrder(node.right, fn)
	fn(node.value)
}

func (t *BinaryTree) Search(value int) bool {
	return t.searchRecursive(t.root, value)
}

func (t *BinaryTree) searchRecursive(parent *Node, value int) bool {
	if parent == nil {
		return false
	}
	if value < parent.value {
		return t.searchRecursive(parent.left, value)
	}
	if value > parent.value {
		return t.searchRecursive(parent.right, value)
	}
	return true
}

func (t *BinaryTree) Delete(value int) {

}

func (t *BinaryTree) deleteRecursive(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if value < node.value {
		node.left = t.deleteRecursive(node.left, value)
		return node
	}
	if value > node.value {
		node.right = t.deleteRecursive(node.right, value)
		return node
	}
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left != nil {
		node = node.left
		return node
	}
	if node.right != nil {
		node = node.right
		return node
	}
	newParent := node.right
	newParent.left = node.left
	node = nil

	return newParent
}

func main() {
	tree := &BinaryTree{}
	tree.Insert(8)
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(5)
	tree.Insert(7)

	fmt.Println(tree.Search(5))
}
