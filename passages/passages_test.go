package passages

import (
	"fmt"
	"sort"
	"testing"
)

var (
	INPUT_SIMPLE  = []string{"start-A", "start-b", "A-c", "A-b", "b-d", "A-end", "b-end"}
	INPUT_MEDIUM  = []string{"dc-end", "HN-start", "start-kj", "dc-start", "dc-HN", "LN-dc", "HN-end", "kj-sa", "kj-HN", "kj-dc"}
	INPUT_COMPLEX = []string{"fs-end", "he-DX", "fs-he", "start-DX", "pj-DX", "end-zg", "zg-sl", "zg-pj", "pj-he", "RW-he", "fs-DX", "pj-RW", "zg-RW", "start-pj", "he-WI", "zg-he", "pj-fs", "start-RW"}
)

func TestParseTunnel(t *testing.T) {
	tests := map[string]string{
		"start-end": "start-end",
		"start-a":   "start-a",
		"a-start":   "start-a",
		"a-end":     "a-end",
		"end-a":     "a-end",
		"end-start": "start-end",
		"a-b":       "a-b",
		"b-a":       "b-a",
	}
	for input, expected := range tests {
		tunnel := ParseTunnel(input)
		if expected != tunnel.String() {
			t.Errorf("Expected to find %s, found %s.", expected, tunnel)
		}
	}
}

func TestParsePassage(t *testing.T) {
	passage := ParsePassage("start,b,end")
	expected := Passage{ParseTunnel("start-b"), ParseTunnel("b-end")}
	if passage.String() != expected.String() {
		t.Errorf("Expected passage to be (%s), but found (%s).", passage, expected)
	}
}

func TestParseCave(t *testing.T) {
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
		cave := ParseCave([]string{input})
		if expected != len(cave) {
			t.Errorf("Expected to find %d tunnels, found %d.", expected, len(cave))
		}
	}
}

func TestIsVisitAlowedSimpleYes(t *testing.T) {
	passage := ParsePassage("start,a")
	tunnel := ParseTunnel("a-b")
	expected := true
	actual := passage.isVisitAllowed(tunnel, false)
	if expected != actual {
		t.Errorf("Expected tunnel %s in passage %s: %v, found: %v.", tunnel, passage, expected, actual)
	}
}

func TestIsVisitAlowedSimpleNo(t *testing.T) {
	passage := ParsePassage("start,a,A")
	tunnel := ParseTunnel("A-a")
	expected := false
	actual := passage.isVisitAllowed(tunnel, false)
	if expected != actual {
		t.Errorf("Expected tunnel %s in passage %s: %v, found: %v.", tunnel, passage, expected, actual)
	}
}

func TestIsVisitAlowedComplexYes(t *testing.T) {
	passage := ParsePassage("start,a,A,b,B")
	tunnel := ParseTunnel("B-a")
	expected := true
	actual := passage.isVisitAllowed(tunnel, true)
	if expected != actual {
		t.Errorf("Expected tunnel %s in passage %s: %v, found: %v.", tunnel, passage, expected, actual)
	}
}

func TestIsVisitAlowedComplexNo(t *testing.T) {
	passage := ParsePassage("start,a,A,a,b,B")
	tunnel := ParseTunnel("B-a")
	expected := false
	actual := passage.isVisitAllowed(tunnel, true)
	if expected != actual {
		t.Errorf("Expected tunnel %s in passage %s: %v, found: %v.", tunnel, passage, expected, actual)
	}
}

func TestLeadsToBigCave(t *testing.T) {
	inputs := map[string]bool{
		"start-A": true,
		"start-b": false,
		"A-c":     false,
		"A-b":     false,
		"b-d":     false,
		"A-end":   false,
		"b-end":   false,
		"BB-end":  false,
		"a-AA":    true,
		"end-AA":  false,
	}
	for input, expected := range inputs {
		actual := ParseTunnel(input).leadsToBigCave()
		if expected != actual {
			t.Errorf("Expected %v for %s, but found %v", expected, input, actual)
		}
	}
}

func TestGetNextTunnels(t *testing.T) {
	cave := ParseCave(INPUT_SIMPLE)
	tunnel := ParseTunnel("start-A")
	expected := []Tunnel{ParseTunnel("A-c"), ParseTunnel("A-b"), ParseTunnel("A-end")}
	actual := cave.getNextTunnels(tunnel)
	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", actual) {
		t.Errorf("Expected to find %v, but found %v.", expected, actual)
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
	for _, passage := range ParseCave(INPUT_SIMPLE).findPassages(Passage{ParseTunnel("start-A")}, false) {
		actuals = append(actuals, passage.String())
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
	for _, passage := range ParseCave(INPUT_SIMPLE).findPassages(Passage{ParseTunnel("start-A")}, true) {
		actuals = append(actuals, passage.String())
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
	actual := ParseCave(INPUT_SIMPLE).CountPassagesOut(false)
	if expected != actual {
		t.Errorf("Expected to find %d passages out, found %d.", expected, actual)
	}
}

func TestCountMorePassagesOutSimple(t *testing.T) {
	expected := 36
	actual := ParseCave(INPUT_SIMPLE).CountPassagesOut(true)
	if expected != actual {
		t.Errorf("Expected to find %d passages out, found %d.", expected, actual)
	}
}

func TestCountPassagesOutMedium(t *testing.T) {
	expected := 19
	actual := ParseCave(INPUT_MEDIUM).CountPassagesOut(false)
	if expected != actual {
		t.Errorf("Expected to find %d passages out, found %d.", expected, actual)
	}
}

func TestCountMorePassagesOutMedium(t *testing.T) {
	expected := 103
	actual := ParseCave(INPUT_MEDIUM).CountPassagesOut(true)
	if expected != actual {
		t.Errorf("Expected to find %d passages out, found %d.", expected, actual)
	}
}

func TestCountPassagesOutComplex(t *testing.T) {
	expected := 226
	actual := ParseCave(INPUT_COMPLEX).CountPassagesOut(false)
	if expected != actual {
		t.Errorf("Expected to find %d passages out, found %d.", expected, actual)
	}
}

func TestCountMorePassagesOutComplex(t *testing.T) {
	expected := 3509
	actual := ParseCave(INPUT_COMPLEX).CountPassagesOut(true)
	if expected != actual {
		t.Errorf("Expected to find %d passages out, found %d.", expected, actual)
	}
}
