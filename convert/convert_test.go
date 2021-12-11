package convert

import (
	"testing"

	"github.com/sneils/adventofcode2021/ints"
)

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected convert to panic, but did not?!")
		}
	}()
	ToIntBase("xxx", 10)
}

func TestToIntBase(t *testing.T) {
	input := "101"
	base := 10
	expected := 101
	actual := ToIntBase(input, base)
	if actual != expected {
		t.Errorf("Expected to get %d, but got %d instead!", expected, actual)
	}
}

func TestToIntsBase(t *testing.T) {
	input := []string{"101", "202"}
	base := 10
	expected := 303
	actual := ints.GetSum(ToIntsBase(input, base))
	if actual != expected {
		t.Errorf("Expected to get %d, but got %d instead!", expected, actual)
	}
}

func TestToInt(t *testing.T) {
	input := "101"
	expected := 5
	actual := FromBinary(input)
	if actual != expected {
		t.Errorf("Expected to get %d, but got %d instead!", expected, actual)
	}
}

func TestToInts(t *testing.T) {
	input := []string{"101", "202"}
	expected := 303
	actual := ints.GetSum(ToInts(input))
	if actual != expected {
		t.Errorf("Expected to get %d, but got %d instead!", expected, actual)
	}
}

func TestFromBinary(t *testing.T) {
	input := "101"
	expected := 5
	actual := FromBinary(input)
	if actual != expected {
		t.Errorf("Expected to get %d, but got %d instead!", expected, actual)
	}
}

func TestFromHex(t *testing.T) {
	input := "ff"
	expected := 255
	actual := FromHex(input)
	if actual != expected {
		t.Errorf("Expected to get %d, but got %d instead!", expected, actual)
	}
}
