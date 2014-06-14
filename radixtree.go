package datastructures

import (
	"errors"
	"strings"
)

type node struct {
	left    *node
	right   *node
	parent  *node
	isValue bool
	value   interface{}
}

type BinaryRadixTree struct {
	root   *node
	length int
}

func NewBinaryRadixTree() *BinaryRadixTree {
	return &BinaryRadixTree{root: &node{isValue: false, value: string(-5), left: nil, right: nil, parent: nil}, length: 0}
}

func (t *BinaryRadixTree) Insert(bits string) error {
	isBinary := isBinaryString(bits)

	if !isBinary {
		return errors.New("string is not binary")
	}

	if len(bits) == 0 {
		return errors.New("can't insert empty string")
	}

	err := insertBits(t, bits)
	if err != nil {
		return err
	}
	t.length = t.length + 1
	return nil
}

//write individual bits into the radix tree
func insertBits(tree *BinaryRadixTree, bits string) error {
	bitSlice := strings.Split(bits, "")
	insertBit(bitSlice, tree, tree.root)
	return nil
}

func insertBit(bitSlice []string, tree *BinaryRadixTree, parentNode *node) {
	if len(bitSlice) == 0 {
		return
	}

	bit := bitSlice[0]
	remainingBitSlice := bitSlice[1:len(bitSlice)]
	isAValue := false
	if len(bitSlice) == 1 {
		isAValue = true
	}

	bitNode := &node{left: nil, right: nil, parent: parentNode, isValue: isAValue, value: bit}

	if bit == "1" {
		if parentNode.right != nil {
			bitNode.left = parentNode.right.left
			bitNode.right = parentNode.right.right
			if parentNode.right.isValue == true {
				bitNode.isValue = true
			}
		}
		parentNode.right = bitNode
	} else {
		if parentNode.left != nil {
			bitNode.left = parentNode.left.left
			bitNode.right = parentNode.left.right
			if parentNode.left.isValue == true {
				bitNode.isValue = true
			}
		}
		parentNode.left = bitNode
	}

	insertBit(remainingBitSlice, tree, bitNode)

}

func (t *BinaryRadixTree) Traverse() {
	traverseNode(t.root, "", nil)
}

func traverseNode(parentNode *node, bitString string, processor func(visitedString string)) {
	if parentNode == nil {
		return
	}
	//only if not root
	if parentNode.parent != nil {
		bitString += parentNode.value.(string)
	}

	if parentNode.isValue {
		if processor != nil {
			processor(bitString)
		}

	}
	if parentNode.left != nil {
		traverseNode(parentNode.left, bitString, processor)
	}

	if parentNode.right != nil {
		traverseNode(parentNode.right, bitString, processor)
	}
}

func (t *BinaryRadixTree) Sort() []string {
	stringChan := make(chan string)
	go traverseNode(t.root, "", func(visitedString string) {
		stringChan <- visitedString
	})

	var sortedStrings []string

	for i := 0; i < t.length; i++ {
		bitString := <-stringChan
		sortedStrings = append(sortedStrings, bitString)
	}

	return sortedStrings
}

func isBinaryString(value string) bool {
	processedValue := strings.Replace(strings.Replace(value, "0", "", -1), "1", "", -1)
	if processedValue == "" {
		return true
	}
	return false
}
