package snailfish

import "testing"

func TestReverse(t *testing.T) {
	tests := map[string]string{
		"3,":                 ",3",
		"[[3,":               ",3[[",
		"][][][][[[]]]]]33,": ",33]]]]][[[][][][]",
	}
	for input, expected := range tests {
		actual := reverse(input)
		if expected != actual {
			t.Errorf("Expected to find %s, found %s.", expected, actual)
		}
	}
}

func TestGetNextNumber(t *testing.T) {
	tests := map[string]string{
		"3,":                 "3",
		"[[3,":               "3",
		"][][][][[[]]]]]33,": "33",
		",":                  "",
	}
	for input, expected := range tests {
		actual := getNextNumber(input)
		if expected != actual {
			t.Errorf("Expected to find %s, found %s.", expected, actual)
		}
	}
}

func TestGetLastNumber(t *testing.T) {
	tests := map[string]string{
		"2,2,3,3,":             "3",
		"3323[[3,":             "3",
		"55][][][][[[]]]]]33,": "33",
		",":                    "",
	}
	for input, expected := range tests {
		actual := getLastNumber(input)
		if expected != actual {
			t.Errorf("Expected to find %s, found %s.", expected, actual)
		}
	}
}

func TestSplit1(t *testing.T) {
	tests := map[int]string{
		10: "[5,5]",
		11: "[5,6]",
		12: "[6,6]",
	}
	for input, expected := range tests {
		actual := Split(input)
		if expected != actual.String() {
			t.Errorf("Expected split to look like %s, but looks like %s.", expected, actual)
		}
	}
}
