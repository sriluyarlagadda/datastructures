package datastructures

import (
	"errors"
	"fmt"
)

var _ = errors.New

//any type implementing this interface should have a compare method
type Comparer interface {
	Compare(y Comparer) (bool, error)
}

type Node struct {
	NodeValue Comparer
	parent    *Node
	left      *Node
	right     *Node
}

func (n *Node) DisplayNode() {
	if n == nil {
		fmt.Println("node does not exist")
	}
	fmt.Println("node value:", n.NodeValue)

	if n.parent != nil {
		fmt.Println("node parent:", n.parent.NodeValue)
	}

	if n.left != nil {
		fmt.Println("node left chid:", n.left.NodeValue)
	}
	if n.right != nil {
		fmt.Println("node right chid:", n.right.NodeValue)
	}
}

type Tree struct {
	head *Node
}

func NewTree() *Tree {
	return &Tree{head: nil}
}

func (t *Tree) Insert(value Comparer) {
	level := 0
	fmt.Println("inserting value:", value)
	if t.head == nil {
		node := &Node{NodeValue: value, left: nil, right: nil}
		t.head = node
		return
	}

	currenetNode := t.head
	var leaf *Node
	for currenetNode != nil {
		leaf = currenetNode

		if comparisonValue, error := currenetNode.NodeValue.Compare(value); error != nil {
			return
		} else {
			if comparisonValue == true {
				currenetNode = currenetNode.left
			} else {
				currenetNode = currenetNode.right
			}
			level++
		}

	}
	node := &Node{NodeValue: value, left: nil, right: nil}
	if comparisonValue, error := leaf.NodeValue.Compare(value); error != nil {
		return
	} else {

		if comparisonValue == true {
			leaf.left = node
		} else {
			leaf.right = node
		}
	}

	node.parent = leaf
}

func (t *Tree) Search(value Comparer) *Node {
	if t == nil {
		fmt.Println("tree empty")
	}
	fmt.Println("searching for...", value)

	nodeChannel := make(chan *Node)

	DoesExist := func(node *Node) {
		if node.NodeValue == value {
			nodeChannel <- node
		}
	}

	traverseStatus := make(chan bool)
	go traverseTree(t, DoesExist, traverseStatus)

	select {
	case val := <-traverseStatus:
		if val == true {
			fmt.Println("search done! ", value, " not found")
		}
		return nil

	case selectedNode := <-nodeChannel:
		return selectedNode
	}

}

func (t *Tree) Traverse() {
	printNode := func(node *Node) {
		fmt.Println("node:", node.NodeValue)

	}
	traverseTree(t, printNode, nil)
}

func traverseTree(t *Tree, f func(node *Node), traverseStatus chan bool) {
	if t.head == nil {
		fmt.Println("tree empty")
	}
	fmt.Println("traversing tree....")

	traverse(t.head, f)

	if traverseStatus != nil {
		traverseStatus <- true
	}
}

func traverse(t *Node, f func(node *Node)) {
	if t == nil {
		return
	}

	if t != nil {
		traverse(t.left, f)
		process(t, f)
		traverse(t.right, f)

	}
}

func process(value *Node, f func(node *Node)) {
	f(value)
}

func (t *Tree) Minimum(node *Node) *Node {
	currentNode := node
	min := currentNode
	if currentNode.left != nil {
		min = currentNode.left
		currentNode = currentNode.left
	}
	return min
}

func (t *Tree) Succeser(node *Node) *Node {
	if node.right != nil {
		return t.Minimum(node.right)
	}
	parent := node.parent
	currentNode := node
	if parent != nil && currentNode == parent.right {
		currentNode = parent
		parent = parent.parent
	}
	return parent
}

func (t *Tree) Transplant(node1, node2 *Node) {
	node1Parent := node1.parent
	if node1Parent == nil {
		t.head = node2
		return
	}

	if node2 == nil {

		if node1Parent.left == node1 {
			node1Parent.left = node2
		} else {
			node1Parent.right = node2
		}
		return
	}

	node2Parent := node2.parent
	if node2Parent.left == node2 {
		node2Parent.left = nil
	} else {
		node2Parent.right = nil
	}

	node2.parent = node1Parent

	if node1Parent.left == node1 {
		node1Parent.left = node2
	} else {
		node1Parent.right = node2
	}

}

func (t *Tree) Delete(node *Node) {

	//if it does not have children
	if node.left == nil && node.right == nil {
		nodeParent := node.parent
		if nodeParent.left == node {
			nodeParent.left = nil
		} else {
			nodeParent.right = nil
		}
		return
	}

	//if it has one child at left
	if node.left != nil && node.right == nil {
		t.Transplant(node, node.left)
		return
	}

	if node.right != nil && node.left == nil {
		t.Transplant(node, node.right)
		return
	}

	if node.left != nil && node.right != nil {
		succeser := t.Succeser(node)
		if node.right != succeser {
			t.Transplant(succeser, succeser.right)
			succeser.right = node.right
			succeser.right.parent = succeser
		}
		t.Transplant(node, succeser)
		succeser.left = node.left
		node.left.parent = succeser

	}

}
