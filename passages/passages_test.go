package passages

import (
	"sort"
	"testing"
)

var (
	INPUT_SIMPLE  = []string{"start-A", "start-b", "A-c", "A-b", "b-d", "A-end", "b-end"}
	INPUT_MEDIUM  = []string{"dc-end", "HN-start", "start-kj", "dc-start", "dc-HN", "LN-dc", "HN-end", "kj-sa", "kj-HN", "kj-dc"}
	INPUT_COMPLEX = []string{"fs-end", "he-DX", "fs-he", "start-DX", "pj-DX", "end-zg", "zg-sl", "zg-pj", "pj-he", "RW-he", "fs-DX", "pj-RW", "zg-RW", "start-pj", "he-WI", "zg-he", "pj-fs", "start-RW"}
)

func TestNewCaveSystem(t *testing.T) {
	tests := map[string]int{
		"start-end": 1,
		"start-a":   1,
		"a-start":   1,
		"a-end":     1,
		"end-a":     1,
		"end-start": 1,
		"a-b":       2,
		"b-a":       2,
	}
	for input, expected := range tests {
		cave := NewCaveSystem([]string{input})
		if expected != len(cave) {
			t.Errorf("Expected to find %d tunnels, found %d.", expected, len(cave))
		}
	}
}

func TestIsVisitAllowed(t *testing.T) {
	tests := map[Passage]bool{
		NewPassage("start,b", false):        true,
		NewPassage("start,a,A", false):      false,
		NewPassage("start,a,A,b,B", false):  false,
		NewPassage("start,a,A,a,b,B", true): false,
	}
	next := "a"
	for passage, expected := range tests {
		actual := passage.isVisitAllowed(next, false)
		if expected != actual {
			t.Errorf("Expected %s next in passage %v: %v, found: %v.", next, passage, expected, actual)
		}
	}
}

func TestIsVisitAllowedRevisit(t *testing.T) {
	tests := map[Passage]bool{
		NewPassage("start,b", false):        true,
		NewPassage("start,a,A", false):      true,
		NewPassage("start,a,A,b,B", false):  true,
		NewPassage("start,a,A,a,b,B", true): false,
	}
	next := "a"
	for passage, expected := range tests {
		actual := passage.isVisitAllowed(next, true)
		if expected != actual {
			t.Errorf("Expected %s next in passage %v: %v, found: %v.", next, passage, expected, actual)
		}
	}
}

func TestFindPassages(t *testing.T) {
	expected := []string{
		"start,A,b,A,c,A,end",
		"start,A,b,A,end",
		"start,A,b,end",
		"start,A,c,A,b,A,end",
		"start,A,c,A,b,end",
		"start,A,c,A,end",
		"start,A,end",
	}
	actuals := []string{}
	for _, passage := range NewCaveSystem(INPUT_SIMPLE).findPassages(NewPassage("start,A", false), false) {
		actuals = append(actuals, passage.visited)
	}
	sort.Strings(actuals)
	if len(expected) != len(actuals) {
		t.Errorf("Expected to find %d passages, found %d", len(expected), len(actuals))
	}
	for idx, actual := range actuals {
		exp := expected[idx]
		if exp != actual {
			t.Errorf("Expected to find passage %s, but found %s.", exp, actual)
		}
	}
}

func TestFindMorePassages(t *testing.T) {
	expected := []string{
		"start,A,b,A,b,A,c,A,end",
		"start,A,b,A,b,A,end",
		"start,A,b,A,b,end",
		"start,A,b,A,c,A,b,A,end",
		"start,A,b,A,c,A,b,end",
		"start,A,b,A,c,A,c,A,end",
		"start,A,b,A,c,A,end",
		"start,A,b,A,end",
		"start,A,b,d,b,A,c,A,end",
		"start,A,b,d,b,A,end",
		"start,A,b,d,b,end",
		"start,A,b,end",
		"start,A,c,A,b,A,b,A,end",
		"start,A,c,A,b,A,b,end",
		"start,A,c,A,b,A,c,A,end",
		"start,A,c,A,b,A,end",
		"start,A,c,A,b,d,b,A,end",
		"start,A,c,A,b,d,b,end",
		"start,A,c,A,b,end",
		"start,A,c,A,c,A,b,A,end",
		"start,A,c,A,c,A,b,end",
		"start,A,c,A,c,A,end",
		"start,A,c,A,end",
		"start,A,end",
	}
	actuals := []string{}
	for _, passage := range NewCaveSystem(INPUT_SIMPLE).findPassages(NewPassage("start,A", false), true) {
		actuals = append(actuals, passage.visited)
	}
	sort.Strings(actuals)
	if len(expected) != len(actuals) {
		t.Errorf("Expected to find %d passages, found %d", len(expected), len(actuals))
	}
	for idx, actual := range actuals {
		exp := expected[idx]
		if exp != actual {
			t.Errorf("Expected to find passage %s, but found %s.", exp, actual)
		}
	}
}

func TestCountPassagesOutSimple(t *testing.T) {
	expected := 10
	actual := NewCaveSystem(INPUT_SIMPLE).CountPassagesOut(false)
	if expected != actual {
		t.Errorf("Expected to find %d passages out, found %d.", expected, actual)
	}
}

func TestCountMorePassagesOutSimple(t *testing.T) {
	expected := 36
	actual := NewCaveSystem(INPUT_SIMPLE).CountPassagesOut(true)
	if expected != actual {
		t.Errorf("Expected to find %d passages out, found %d.", expected, actual)
	}
}

func TestCountPassagesOutMedium(t *testing.T) {
	expected := 19
	actual := NewCaveSystem(INPUT_MEDIUM).CountPassagesOut(false)
	if expected != actual {
		t.Errorf("Expected to find %d passages out, found %d.", expected, actual)
	}
}

func TestCountMorePassagesOutMedium(t *testing.T) {
	expected := 103
	actual := NewCaveSystem(INPUT_MEDIUM).CountPassagesOut(true)
	if expected != actual {
		t.Errorf("Expected to find %d passages out, found %d.", expected, actual)
	}
}

func TestCountPassagesOutComplex(t *testing.T) {
	expected := 226
	actual := NewCaveSystem(INPUT_COMPLEX).CountPassagesOut(false)
	if expected != actual {
		t.Errorf("Expected to find %d passages out, found %d.", expected, actual)
	}
}

func TestCountMorePassagesOutComplex(t *testing.T) {
	expected := 3509
	actual := NewCaveSystem(INPUT_COMPLEX).CountPassagesOut(true)
	if expected != actual {
		t.Errorf("Expected to find %d passages out, found %d.", expected, actual)
	}
}
