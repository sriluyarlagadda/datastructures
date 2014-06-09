package datastructures

import (
	"strings"
)

type node struct {
	left    *node
	right   *node
	isValue bool
	value   interface{}
}

type BinaryRadixTree struct {
	root *node
}

func NewBinaryRadixTree() *BinaryRadixTree {
	return &BinaryRadixTree{root: nil}
}

/*func (t *BinaryRadixTree) Insert(value string) error {
	processedValue := strings.Replace(strings.Replace(value, "0", "", -1), "1", "", -1)
}*/

func isBinaryString(value string) bool {
	processedValue := strings.Replace(strings.Replace(value, "0", "", -1), "1", "", -1)
	if processedValue == "" {
		return true
	}
	return false
}
