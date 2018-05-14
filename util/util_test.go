package util

import (
	"testing"
)

func TestCanAddNumbers(t *testing.T) {
	result := Add(1, 2)

	if result != 3 {
		t.Log("Failed to add two to one")
		t.Fail()
	}
}

func TextCanMultipleTwoNumbers(t *testing.T) {
	actual := Multiple(3, 2)
	if actual != 6 {
		t.Log("failed to multiple 3 by 2")
		t.Fail()
	}
}
