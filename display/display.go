package display

import (
	"sort"
	"strings"

	"github.com/sneils/adventofcode2021/convert"
)

type Display struct {
	patterns, outputs []string
}

func ParseString(input string) *Display {
	display := &Display{}
	splits := strings.Split(input, " | ")
	display.patterns = strings.Fields(splits[0])
	display.outputs = strings.Fields(splits[1])
	return display
}

func (display *Display) CountUniqueOutputForms() int {
	count := 0
	for _, output := range display.outputs {
		l := len(output)
		if l == 2 || l == 3 || l == 4 || l == 7 {
			count += 1
		}
	}
	return count
}

func sortCharacters(input string) string {
	runes := []rune(input)
	sort.Slice(runes, func(a, b int) bool {
		return runes[a] < runes[b]
	})
	return string(runes)
}

func removeCharacters(input, characters string) string {
	result := input
	for _, remove := range characters {
		result = strings.Replace(result, string(remove), "", 1)
	}
	return result
}

func (display *Display) getPatternsWithLength(length int) []string {
	patterns := []string{}
	for _, pattern := range display.patterns {
		if len(pattern) == length {
			patterns = append(patterns, pattern)
		}
	}
	return patterns
}

func (display *Display) DeduceLetters() map[string]string {
	numbers := map[string]string{}

	// unique length patterns
	numbers["1"] = display.getPatternsWithLength(2)[0]
	numbers["4"] = display.getPatternsWithLength(4)[0]
	numbers["7"] = display.getPatternsWithLength(3)[0]
	numbers["8"] = display.getPatternsWithLength(7)[0]

	// length 5: 2 3 5
	for _, pattern := range display.getPatternsWithLength(5) {
		check2 := removeCharacters(pattern, numbers["4"])
		if len(check2) == 3 {
			numbers["2"] = pattern
		}

		check3 := removeCharacters(pattern, numbers["7"])
		if len(check3) == 2 {
			numbers["3"] = pattern
		}
	}
	for _, pattern := range display.getPatternsWithLength(5) {
		check5 := removeCharacters(pattern, numbers["2"])
		if len(check5) == 2 {
			numbers["5"] = pattern
		}
	}
	// length 6: 0 6 9
	for _, pattern := range display.getPatternsWithLength(6) {
		check0 := removeCharacters(pattern, numbers["5"])
		if len(check0) == 2 {
			numbers["0"] = pattern
		}

		check6 := removeCharacters(pattern, numbers["7"])
		if len(check6) == 4 {
			numbers["6"] = pattern
		}

		check9 := removeCharacters(pattern, numbers["3"])
		if len(check9) == 1 {
			numbers["9"] = pattern
		}
	}

	return map[string]string{
		removeCharacters(numbers["7"], numbers["1"]):                                 "a",
		removeCharacters(numbers["4"], numbers["3"]):                                 "b",
		removeCharacters(numbers["8"], numbers["6"]):                                 "c",
		removeCharacters(numbers["8"], numbers["0"]):                                 "d",
		removeCharacters(numbers["8"], numbers["9"]):                                 "e",
		removeCharacters(numbers["7"], numbers["2"]):                                 "f",
		removeCharacters(removeCharacters(numbers["3"], numbers["4"]), numbers["7"]): "g",
	}
}

func (display *Display) SumOutputs() int {
	config := map[string]string{
		"abcefg":  "0",
		"cf":      "1",
		"acdeg":   "2",
		"acdfg":   "3",
		"bcdf":    "4",
		"abdfg":   "5",
		"abdefg":  "6",
		"acf":     "7",
		"abcdefg": "8",
		"abcdfg":  "9",
	}
	letters := display.DeduceLetters()
	num := ""
	for _, output := range display.outputs {
		converted := ""
		for _, char := range output {
			converted += letters[string(char)]
		}
		sorted := sortCharacters(converted)
		num += config[sorted]
	}
	return convert.ToInt(num)
}
