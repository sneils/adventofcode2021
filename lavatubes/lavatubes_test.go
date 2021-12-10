package lavatubes

import (
	"fmt"
	"testing"
)

var SAMPLE_INPUT = []string{
	"2199943210",
	"3987894921",
	"9856789892",
	"8767896789",
	"9899965678",
}

func TestGetNeighbors(t *testing.T) {
	tubes := Parse(SAMPLE_INPUT)
	expected := []LavaTube{
		{1, 0, 1},
		{0, 1, 3},
		{2, 1, 8},
		{1, 2, 8},
	}
	found := tubes.getNeighbors(LavaTube{1, 1, 9})
	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", found) {
		t.Errorf("Expected to find %v as neighbors, found %v.", expected, found)
	}
}

func TestGetLowpoints(t *testing.T) {
	tubes := Parse(SAMPLE_INPUT)
	expected := []LavaTube{
		{1, 0, 1},
		{9, 0, 0},
		{2, 2, 5},
		{6, 4, 5},
	}
	found := tubes.getLowpoints()
	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", found) {
		t.Errorf("Expected to find %v lowpoints, found %v.", expected, found)
	}
}

func TestCalculateRisk(t *testing.T) {
	tubes := Parse(SAMPLE_INPUT)
	expected := 15
	found := tubes.CalculateRisk()
	if expected != found {
		t.Errorf("Expected risk to be %d, found %d.", expected, found)
	}
}

func TestContains(t *testing.T) {
	tubes := []LavaTube{
		{2, 2, 0},
		{3, 3, 0},
		{1, 1, 0},
	}
	find := LavaTube{1, 1, 0}
	expected := true
	found := contains(tubes, find)
	if expected != found {
		t.Errorf("Expected to find %v in %v, but didn't.", find, tubes)
	}
}

func TestGetBasin(t *testing.T) {
	tubes := Parse(SAMPLE_INPUT)
	expected := []LavaTube{
		{1, 0, 1},
		{0, 0, 2},
		{0, 1, 3},
	}
	lowpoint := LavaTube{1, 0, 1}
	found := tubes.getBasin(lowpoint, []LavaTube{})
	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", found) {
		t.Errorf("Expected to find basin %v, found %v.", expected, found)
	}
}

func TestGetBiggestThreeBasinSizes(t *testing.T) {
	tubes := Parse(SAMPLE_INPUT)
	expected := []int{9, 9, 14}
	found := tubes.GetBiggestThreeBasinSizes()
	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", found) {
		t.Errorf("Expected to find sizes %v, found %v.", expected, found)
	}
}
