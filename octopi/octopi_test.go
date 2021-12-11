package octopi

import (
	"log"
	"strings"
	"testing"
)

var SAMPLE_INPUT = []string{
	"5483143223",
	"2745854711",
	"5264556173",
	"6141336146",
	"6357385478",
	"4167524645",
	"2176841721",
	"6882881134",
	"4846848554",
	"5283751526",
}

func TestEnergizeNeighbors(t *testing.T) {
	grid := Parse(SAMPLE_INPUT)
	grid.energizeNeighbors(Octopus{1, 1, 7})
	expected := []string{
		"6593143223",
		"3755854711",
		"6374556173",
		"6141336146",
		"6357385478",
		"4167524645",
		"2176841721",
		"6882881134",
		"4846848554",
		"5283751526",
	}
	actual := grid.Pack()
	if strings.Join(expected, ",") != strings.Join(actual, ",") {
		t.Errorf("Didn't meet expectations :(")
		log.Println("Expected:")
		printStringSlice(expected)
		log.Println("Actual:")
		printStringSlice(actual)
	}
}

func TestContains(t *testing.T) {
	octopi := []Octopus{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
	}
	find := Octopus{2, 2, 2}
	dont := Octopus{4, 4, 4}
	if !contains(octopi, find) {
		t.Errorf("Didn't find %v in %v :(", find, octopi)
	}
	if contains(octopi, dont) {
		t.Errorf("Did find %v in %d :(", dont, octopi)
	}
}

func TestPack(t *testing.T) {
	grid := Parse(SAMPLE_INPUT)
	expected := SAMPLE_INPUT
	actual := grid.Pack()
	if strings.Join(expected, ",") != strings.Join(actual, ",") {
		t.Errorf("Didn't meet expectations :(")
		log.Println("Expected:")
		printStringSlice(expected)
		log.Println("Actual:")
		printStringSlice(actual)
	}
}

func TestDoAge100(t *testing.T) {
	grid := Parse(SAMPLE_INPUT)
	for grid.Age < 100 {
		grid.DoAge()
	}
	expectedAge := 100
	expectedFlashes := 1656
	expectedGrid := []string{
		"0397666866",
		"0749766918",
		"0053976933",
		"0004297822",
		"0004229892",
		"0053222877",
		"0532222966",
		"9322228966",
		"7922286866",
		"6789998766",
	}
	actualAge := grid.Age
	actualFlashes := grid.Flashes
	actualGrid := grid.Pack()
	if strings.Join(expectedGrid, ",") != strings.Join(actualGrid, ",") {
		t.Errorf("Didn't meet expectations :(")
		log.Println("Expected:")
		printStringSlice(expectedGrid)
		log.Println("Actual:")
		printStringSlice(actualGrid)
	}
	if expectedAge != actualAge {
		t.Errorf("Expected age to be %d, found it to be %d.", expectedAge, actualAge)
	}
	if expectedFlashes != actualFlashes {
		t.Errorf("Expected to find %d flashes, found %d.", expectedFlashes, actualFlashes)
	}
}

func TestForwardToFirstFullFLash(t *testing.T) {
	grid := Parse(SAMPLE_INPUT)
	grid.ForwardToFirstFullFlash()
	expected := 195
	actual := grid.Age
	if expected != actual {
		t.Errorf("Expected age to be %d, found it to be %d.", expected, actual)
	}
}
