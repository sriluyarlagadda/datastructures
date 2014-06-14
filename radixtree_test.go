package datastructures

import (
	"fmt"
	"strconv"
	"testing"
)

var tests map[string]bool = map[string]bool{
	"1011":  true,
	"abc":   false,
	"a  c":  false,
	"1a2z":  false,
	"00193": false,
	"000":   true,
}

func TestIsBinaryString(t *testing.T) {

	for inputString, expectedValue := range tests {
		if isBinary := isBinaryString(inputString); isBinary != expectedValue {
			t.Errorf("expected ", expectedValue, " but returns", isBinary)
		}
	}
}

func TestInsert(t *testing.T) {
	radixTree := NewBinaryRadixTree()
	err := radixTree.Insert("1a00")
	if err == nil || err.Error() != "string is not binary" {
		t.Errorf("expected", err.Error(), " but returns", err.Error())
	}

	err = radixTree.Insert("")
	if err == nil || err.Error() != "can't insert empty string" {
		t.Errorf("expected", err.Error(), " but returns", err.Error())
	}
}

func TestInsertString(t *testing.T) {
	radixTree := NewBinaryRadixTree()
	radixTree.Insert("101")
	fmt.Println(radixTree)
	if radixTree.root.right.value != "1" {
		t.Errorf("root: expected", 1, "actual", radixTree.root.value)
	}

	if radixTree.root.right.left.value != "0" {
		t.Errorf("leftchild: expected", 0, "actual", radixTree.root.left.value)

	}

	if radixTree.root.right.right != nil {
		t.Errorf("rightchild: expected", "nil", radixTree.root.right, radixTree.root.right)
	}

	if radixTree.root.right.left.isValue == true {
		t.Errorf("isIntermediatevalue: expected", false, "actual", strconv.FormatBool(radixTree.root.left.isValue))
	}

	if radixTree.root.right.left.right.isValue == false {
		t.Errorf("isIntermediatevalue: expected", true, "actual", strconv.FormatBool(radixTree.root.left.right.isValue))

	}
}
