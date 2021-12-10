package lanternfish

import (
	"testing"
)

func TestParseString(t *testing.T) {
	inputs := "3,4"
	expected := map[int]int{3: 1, 4: 1}
	school := ParseString(inputs)

	for key, val := range expected {
		got := school.Get(key)
		if got != val {
			t.Errorf("Expected school to have %d fish that are %d years old, found %d instead.", val, key, got)
		}
	}
}

func TestDoAge(t *testing.T) {
	expected := []int{
		5,  // "2,3,2,0,1",
		6,  // "1,2,1,6,0,8",
		7,  // "0,1,0,5,6,7,8",
		9,  // "6,0,6,4,5,6,7,8,8",
		10, // "5,6,5,3,4,5,6,7,7,8",
		10, // "4,5,4,2,3,4,5,6,6,7",
		10, // "3,4,3,1,2,3,4,5,5,6",
		10, // "2,3,2,0,1,2,3,4,4,5",
		11, // "1,2,1,6,0,1,2,3,3,4,8",
		12, // "0,1,0,5,6,0,1,2,2,3,7,8",
		15, // "6,0,6,4,5,6,0,1,1,2,6,7,8,8,8",
		17, // "5,6,5,3,4,5,6,0,0,1,5,6,7,7,7,8,8",
		19, // "4,5,4,2,3,4,5,6,6,0,4,5,6,6,6,7,7,8,8",
		20, // "3,4,3,1,2,3,4,5,5,6,3,4,5,5,5,6,6,7,7,8",
		20, // "2,3,2,0,1,2,3,4,4,5,2,3,4,4,4,5,5,6,6,7",
		21, // "1,2,1,6,0,1,2,3,3,4,1,2,3,3,3,4,4,5,5,6,8",
		22, // "0,1,0,5,6,0,1,2,2,3,0,1,2,2,2,3,3,4,4,5,7,8",
		26, // "6,0,6,4,5,6,0,1,1,2,6,0,1,1,1,2,2,3,3,4,6,7,8,8,8,8",
	}
	inputs := "3,4,3,1,2"
	school := ParseString(inputs)
	for day, ex := range expected {
		school.DoAge()
		count := school.Count()
		if count != ex {
			t.Fatalf("Expected school to have %d fish on day %d, found %d.", ex, day+1, count)
		}
	}
}
