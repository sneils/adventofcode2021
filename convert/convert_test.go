package convert

import "testing"

func TestToIntBase(t *testing.T) {
	input := "101"
	base := 10
	expected := 101
	output := ToIntBase(input, base)
	if output != expected {
		t.Errorf("Expected to get %d, but got %d instead!", expected, output)
	}
}

func TestFromHex(t *testing.T) {
	input := "ff"
	expected := 255
	output := FromHex(input)
	if output != expected {
		t.Errorf("Expected to get %d, but got %d instead!", expected, output)
	}
}

func TestToInt(t *testing.T) {
	input := "101"
	expected := 5
	output := FromBinary(input)
	if output != expected {
		t.Errorf("Expected to get %d, but got %d instead!", expected, output)
	}
}

func TestFromBinary(t *testing.T) {
	input := "101"
	expected := 5
	output := FromBinary(input)
	if output != expected {
		t.Errorf("Expected to get %d, but got %d instead!", expected, output)
	}
}
