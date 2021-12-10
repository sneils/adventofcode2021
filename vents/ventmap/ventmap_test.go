package ventmap

import (
	"testing"

	"github.com/sneils/adventofcode2021/vents"
)

func getSampleMap(allowDiagonals bool) *VentMap {
	sample := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}
	vm := New(allowDiagonals)
	for _, input := range sample {
		vent := vents.ParseString(input)
		vm.Add(vent)
	}
	return vm
}

func TestWithoutDiagonals(t *testing.T) {
	expected := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 1, 1, 2, 1, 1, 1, 2, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
	}
	vm := getSampleMap(false)
	for y, row := range expected {
		for x, col := range row {
			if col == 0 {
				continue
			}
			val := vm.Get(x, y)
			if col != val {
				t.Errorf("Expected %d @ %d,%d, found %d.", col, x, y, val)
			}
		}
	}
	score := vm.GetOverlaps()
	expectedScore := 5
	if score != expectedScore {
		t.Errorf("Expected a score of %d, found %d.", expectedScore, score)
	}
}

func TestWithDiagonals(t *testing.T) {
	expected := [][]int{
		{1, 0, 1, 0, 0, 0, 0, 1, 1, 0},
		{0, 1, 1, 1, 0, 0, 0, 2, 0, 0},
		{0, 0, 2, 0, 1, 0, 1, 1, 1, 0},
		{0, 0, 0, 1, 0, 2, 0, 2, 0, 0},
		{0, 1, 1, 2, 3, 1, 3, 2, 1, 1},
		{0, 0, 0, 1, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
	}
	vm := getSampleMap(true)
	for y, row := range expected {
		for x, col := range row {
			if col == 0 {
				continue
			}
			val := vm.Get(x, y)
			if col != val {
				t.Errorf("Expected %d @ %d,%d, found %d.", col, x, y, val)
			}
		}
	}
	score := vm.GetOverlaps()
	expectedScore := 12
	if score != expectedScore {
		t.Errorf("Expected a score of %d, found %d.", expectedScore, score)
	}
}
