package days

import (
	"bufio"
	"fmt"
	"os"
)

type Day struct {
	Inputs []string
}

func New(day int) *Day {
	return &Day{Inputs: readPuzzleFile(day)}
}

func readPuzzleFile(day int) []string {
	path := fmt.Sprintf("./days/input/%d.txt", day)

	open, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(open)
	inputs := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		inputs = append(inputs, text)
	}

	return inputs
}

func (day *Day) Groups() [][]string {
	groups := [][]string{}
	group := []string{}

	for _, input := range day.Inputs {
		if input == "" {
			groups = append(groups, group)
			group = []string{}
			continue
		}
		group = append(group, input)
	}

	return append(groups, group)
}
