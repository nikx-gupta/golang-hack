package trees

import (
	"sync"
)

func main() {

}

type TreeNode struct {
	key       int
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

func (t *BinarySearchTree) InsertElement(key int, value int) {
	t.lock.Lock()
	defer t.lock.Unlock()
	var treeNode *TreeNode
	treeNode = &TreeNode{key, value, nil, nil}
	if t.rootNode == nil {
		t.rootNode = treeNode
	} else {
		insertNode(t.rootNode, treeNode)
	}
}

func (t *BinarySearchTree) InOrderTraverse(function func(int, string)) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	inOrderTraverse(t.rootNode, "Root", function)
}

func (t *BinarySearchTree) PreOrderTraverse(function func(int, string)) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	preOrderTraverse(t.rootNode, "Root", function)
}

func (t *BinarySearchTree) PostOrderTraverse(function func(int, string)) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	postOrderTraverse(t.rootNode, "Root", function)
}

func (t *BinarySearchTree) MinNode() *int {
	var curNode = t.rootNode
	if curNode == nil {
		return (*int)(nil)
	}

	for {
		if curNode.leftNode != nil {
			return &curNode.value
		}

		curNode = curNode.leftNode
	}
}

func (t *BinarySearchTree) MaxNode() *int {
	var curNode *TreeNode = t.rootNode
	if curNode == nil {
		return (*int)(nil)
	}

	for {
		if curNode.rightNode != nil {
			return &curNode.value
		}

		curNode = curNode.rightNode
	}
}

func (t *BinarySearchTree) SearchNode(key int) bool {
	return searchNode(t.rootNode, key)
}

func searchNode(node *TreeNode, key int) bool {
	if node == nil {
		return false
	}

	if key < node.key {
		return searchNode(node.leftNode, key)
	}

	if key > node.key {
		return searchNode(node.rightNode, key)
	}

	return true
}

func postOrderTraverse(node *TreeNode, leftOrR string, function func(int, string)) {
	if node != nil {
		// Left -> Right -> CurrentNode
		inOrderTraverse(node.leftNode, "L", function)
		inOrderTraverse(node.rightNode, "R", function)
		function(node.value, leftOrR)
	}
}

func inOrderTraverse(node *TreeNode, leftOrR string, function func(int, string)) {
	if node != nil {
		// Left -> Current -> Right
		inOrderTraverse(node.leftNode, "L", function)
		function(node.value, leftOrR)
		inOrderTraverse(node.rightNode, "R", function)
	}
}

func preOrderTraverse(node *TreeNode, leftOrR string, function func(int, string)) {
	if node != nil {
		// Current -> Left -> Right
		function(node.value, leftOrR)
		inOrderTraverse(node.leftNode, "L", function)
		inOrderTraverse(node.rightNode, "R", function)
	}
}

func insertNode(curNode *TreeNode, newNode *TreeNode) {
	if newNode.key > curNode.key {
		if curNode.rightNode != nil {
			insertNode(curNode.rightNode, newNode)
		} else {
			curNode.rightNode = newNode
		}
	} else {
		if curNode.leftNode != nil {
			insertNode(curNode.leftNode, newNode)
		} else {
			curNode.leftNode = newNode
		}
	}
}
