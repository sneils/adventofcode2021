package display

import "testing"

var INPUT = "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"

func TestDeduceLetters(t *testing.T) {
	display := ParseString(INPUT)
	expected := map[string]string{
		"d": "a",
		"e": "b",
		"a": "c",
		"f": "d",
		"g": "e",
		"b": "f",
		"c": "g",
	}
	letters := display.DeduceLetters()
	for key, val := range letters {
		if val != expected[key] {
			t.Errorf("Expected that %s resolves to %s, found it resolves to %s.", key, expected[key], val)
		}
	}
}

func TestCountUniqueOutputForms(t *testing.T) {
	display := ParseString(INPUT)
	expected := 0
	found := display.CountUniqueOutputForms()
	if expected != found {
		t.Errorf("Expected to find %d unique output forms, found %d.", expected, found)
	}
}

func TestSumOutput(t *testing.T) {
	display := ParseString(INPUT)
	expected := 5353
	found := display.SumOutputs()
	if expected != found {
		t.Errorf("Expected output sum to be %d, found %d.", expected, found)
	}
}

func TestGetPatternsWithLength(t *testing.T) {
	display := ParseString(INPUT)
	input := 5
	expected := 3
	found := display.getPatternsWithLength(input)
	if expected != len(found) {
		t.Errorf("Expected to find %d patterns, found %d.", expected, len(found))
	}
}

func TestRemoveCharacters(t *testing.T) {
	input := "abcdef"
	remove := "cdf"
	expected := "abe"
	found := removeCharacters(input, remove)
	if expected != found {
		t.Errorf("Expected to find `%s` after removing `%s` from `%s`, found `%s`.", expected, remove, input, found)
	}
}

func TestSortCharacters(t *testing.T) {
	input := "bdac"
	expected := "abcd"
	output := sortCharacters(input)
	if expected != output {
		t.Fatalf("Expected sorted string to be `%s`, found `%s`.", expected, output)
	}
}
