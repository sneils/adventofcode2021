package crabmarines

import "testing"

var SAMPLE_INPUT = "16,1,2,0,4,2,7,1,2,14"

func TestAlignConstant(t *testing.T) {
	crabs := NewCrabmarines(SAMPLE_INPUT)
	expectedPosition, expectedFuel := 2, 37
	position, fuel := crabs.AlignConstant()
	if expectedPosition != position {
		t.Errorf("Expected ideal position to be %d, found it was %d.", expectedPosition, position)
	}
	if expectedFuel != fuel {
		t.Errorf("Expected there to be a fuel usage of %d, found it was %d.", expectedFuel, fuel)
	}
}

func TestAlignIncreasing(t *testing.T) {
	crabs := NewCrabmarines(SAMPLE_INPUT)
	expectedPosition, expectedFuel := 5, 168
	position, fuel := crabs.AlignIncreasing()
	if expectedPosition != position {
		t.Errorf("Expected ideal position to be %d, found it was %d.", expectedPosition, position)
	}
	if expectedFuel != fuel {
		t.Errorf("Expected there to be a fuel usage of %d, found it was %d.", expectedFuel, fuel)
	}
}
