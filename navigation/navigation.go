package navigation

import (
	"sort"
)

var CORRUPT_POINT_VALUES = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}
var FIX_POINT_VALUES = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}
var CHUNK_BORDERS = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func isCorrupted(input string) (bool, rune, []rune) {
	stack, fixes := []rune{rune(input[0])}, []rune{}
	for _, char := range input[1:] {
		if _, ok := CHUNK_BORDERS[char]; ok {
			stack = append(stack, char)
			continue
		}
		last := stack[len(stack)-1]
		if char != CHUNK_BORDERS[last] {
			return true, char, fixes
		}
		stack = stack[:len(stack)-1]
	}
	for _, char := range stack {
		fixes = append(fixes, CHUNK_BORDERS[char])
	}
	for a, b := 0, len(fixes)-1; a < b; a, b = a+1, b-1 {
		fixes[a], fixes[b] = fixes[b], fixes[a]
	}
	return false, 'x', fixes
}

func getFixScore(fixes []rune) int {
	score := 0
	for _, fix := range fixes {
		score *= 5
		score += FIX_POINT_VALUES[fix]
	}
	return score
}

func Analyse(input []string) (int, int) {
	corruptionSum, fixSums := 0, []int{}
	for _, line := range input {
		corrupted, char, fixes := isCorrupted(line)
		if corrupted {
			corruptionSum += CORRUPT_POINT_VALUES[char]
			continue
		}
		fixSums = append(fixSums, getFixScore(fixes))
	}
	sort.Ints(fixSums)
	return corruptionSum, fixSums[len(fixSums)/2]
}
