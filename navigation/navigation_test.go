package navigation

import "testing"

var SAMPLE_INPUT = []string{
	"[({(<(())[]>[[{[]{<()<>>",
	"[(()[<>])]({[<{<<[]>>(",
	"{([(<{}[<>[]}>{[]{[(<()>",
	"(((({<>}<{<{<>}{[]{[]{}",
	"[[<[([]))<([[{}[[()]]]",
	"[{[{({}]{}}([{[{{{}}([]",
	"{<[[]]>}<{[{[{[]{()[[[]",
	"[<(<(<(<{}))><([]([]()",
	"<{([([[(<>()){}]>(<<{{",
	"<{([{{}}[<[[[<>{}]]]>[]]",
}

func TestIsNotCorrupted(t *testing.T) {
	line := SAMPLE_INPUT[0]
	expectedBool, expectedFix := false, []rune("}}]])})]")
	actualBool, _, actualFix := isCorrupted(line)
	if expectedBool != actualBool {
		t.Fatalf("Expected line to be _not_ corrupted, but it was :(")
	}
	if string(expectedFix) != string(actualFix) {
		t.Fatalf("Expected to find fix %s, but found %s instead.", string(expectedFix), string(actualFix))
	}
}

func TestIsCorrupted(t *testing.T) {
	line := SAMPLE_INPUT[len(SAMPLE_INPUT)-2]
	expectedBool, expectedChar, expectedFix := true, '>', []rune{}
	actualBool, actualChar, actualFix := isCorrupted(line)
	if expectedBool != actualBool {
		t.Fatalf("Expected line to be corrupted, but it was not :(")
	}
	if expectedChar != actualChar {
		t.Fatalf("Expected to find character %c, but found %c instead.", expectedChar, actualChar)
	}
	if string(expectedFix) != string(actualFix) {
		t.Fatalf("Expected to find fix %s, but found %s instead.", string(expectedFix), string(actualFix))
	}
}

func TestGetFixScore(t *testing.T) {
	input := []rune("}}]])})]")
	expected := 288957
	actual := getFixScore(input)
	if expected != actual {
		t.Fatalf("Expected score to be %d, but found %d instead.", expected, actual)
	}
}

func TestAnalyse(t *testing.T) {
	expectedCorruption, expectedFixes := 26397, 288957
	actualCorruption, actualFixes := Analyse(SAMPLE_INPUT)
	if expectedCorruption != actualCorruption {
		t.Fatalf("Expected syntax error score to be %d, but found it to be %d.", expectedFixes, actualFixes)
	}
	if expectedFixes != actualFixes {
		t.Fatalf("Expected syntax fix score to be %d, but found it to be %d.", expectedFixes, actualFixes)
	}
}
