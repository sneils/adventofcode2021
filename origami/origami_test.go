package origami

import "testing"

var SAMPLE_INPUT = []string{
	"6,10",
	"0,14",
	"9,10",
	"0,3",
	"10,4",
	"4,11",
	"6,0",
	"6,12",
	"4,1",
	"0,13",
	"10,12",
	"3,4",
	"3,0",
	"8,4",
	"1,10",
	"2,14",
	"8,10",
	"9,0",
}

func TestParseFold(t *testing.T) {
	exDir, exPos := "y", 7
	dir, pos := ParseFold("fold along y=7")
	if exDir != dir {
		t.Errorf("Expected dir to be %s, found %s.", exDir, dir)
	}
	if exPos != pos {
		t.Errorf("Expected pos to be %d, found %d.", exPos, pos)
	}
}

func TestFold(t *testing.T) {
	paper := Parse(SAMPLE_INPUT)
	paper.Fold("fold along y=7")
	expected := 17
	actual := paper.CountDots()
	if expected != actual {
		t.Errorf("Expected to see %d dots, but found %d.", expected, actual)
	}
}
