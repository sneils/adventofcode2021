package vents

import (
	"testing"
)

func TestParseString(t *testing.T) {
	input := "0,9 -> 5,9"
	expected := New(0, 9, 5, 9)
	actual := ParseString(input)
	if expected.String() != actual.String() {
		t.Errorf("Expected (%s), found (%s).", expected, actual)
	}
}
