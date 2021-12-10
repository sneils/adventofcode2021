package crabsubmarines

import "testing"

func getSampleCrabs() *CrabSubmarines {
	inputs := "16,1,2,0,4,2,7,1,2,14"
	return ParseString(inputs)
}

func TestParseString(t *testing.T) {
	crabs := getSampleCrabs()
	expected := 10
	count := crabs.Count()
	if expected != count {
		t.Errorf("Expected there to be %d crabs, found %d.", expected, count)
	}
}

func TestAlignConstant(t *testing.T) {
	crabs := getSampleCrabs()
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
	crabs := getSampleCrabs()
	expectedPosition, expectedFuel := 5, 168
	position, fuel := crabs.AlignIncreasing()
	if expectedPosition != position {
		t.Errorf("Expected ideal position to be %d, found it was %d.", expectedPosition, position)
	}
	if expectedFuel != fuel {
		t.Errorf("Expected there to be a fuel usage of %d, found it was %d.", expectedFuel, fuel)
	}
}
