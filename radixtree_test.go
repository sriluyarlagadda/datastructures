package datastructures

import (
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
