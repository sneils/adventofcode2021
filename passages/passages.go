package passages

import (
	"strings"
)

func ParseTunnel(input string) (string, string) {
	split := strings.Split(input, "-")
	if split[0] == "end" || split[1] == "start" {
		return split[1], split[0]
	}
	return split[0], split[1]
}

type Passage struct {
	visited   string
	revisited bool
}

func NewPassage(visited string, revisited bool) Passage {
	return Passage{visited, revisited}
}

func (passage Passage) isVisitAllowed(out string, allowOneRevisit bool) bool {
	if strings.ToUpper(out) == out {
		return true
	}
	if !strings.Contains(passage.visited, ","+out) {
		return true
	}
	if !allowOneRevisit {
		return false
	}
	return !passage.revisited
}

type CaveSystem map[string][]string

func NewCaveSystem(inputs []string) CaveSystem {
	cave := CaveSystem{}
	for _, input := range inputs {
		in, out := ParseTunnel(input)
		cave[in] = append(cave[in], out)
		if in == "start" || out == "end" {
			continue
		}
		cave[out] = append(cave[out], in)
	}
	return cave
}

func (cave CaveSystem) findPassages(passage Passage, allowSingleRevisit bool) []Passage {
	passages := []Passage{}
	holes := strings.Split(passage.visited, ",")
	for _, next := range cave[holes[len(holes)-1]] {
		out := "," + next
		if next == "end" {
			end := passage
			end.visited += out
			passages = append(passages, end)
			continue
		}
		if !passage.isVisitAllowed(next, allowSingleRevisit) {
			continue
		}
		new := passage
		new.revisited = passage.revisited || (strings.ToLower(out) == out && strings.Contains(passage.visited, out))
		new.visited += out
		passages = append(passages, cave.findPassages(new, allowSingleRevisit)...)
	}
	return passages
}

func (cave CaveSystem) CountPassagesOut(allowSingleRevisit bool) int {
	count := 0
	for _, out := range cave["start"] {
		passage := NewPassage("start,"+out, false)
		count += len(cave.findPassages(passage, allowSingleRevisit))

	}
	return count
}
