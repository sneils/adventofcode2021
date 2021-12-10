package submarines

import (
	"strings"

	"github.com/sneils/adventofcode2021/convert"
)

type Submarine struct {
	Horizontal, Depth                        int
	GammaRate, EpsilonRate                   string
	OxygenGeneratorRating, CO2ScrubberRating string
}

func New() *Submarine {
	return &Submarine{0, 0, "", "", "", ""}
}

func (sub *Submarine) Run(cmd Command) {
	if cmd.Name == "up" {
		sub.Depth -= cmd.Value
	} else if cmd.Name == "down" {
		sub.Depth += cmd.Value
	} else if cmd.Name == "forward" {
		sub.Horizontal += cmd.Value
	} else {
		panic("Unknown command :(")
	}
}

func (sub *Submarine) GetPosition() int {
	return sub.Horizontal * sub.Depth
}

func (sub *Submarine) GetPowerConsumption() int {
	return convert.FromBinary(sub.GammaRate) * convert.FromBinary(sub.EpsilonRate)
}

func (sub *Submarine) GetLifeSupportRating() int {
	return convert.FromBinary(sub.CO2ScrubberRating) * convert.FromBinary(sub.OxygenGeneratorRating)
}

func getMostCommonBit(input []string, idx int, defaultBit rune) rune {
	count := 0
	for _, line := range input {
		if line[idx] == '0' {
			count -= 1
		} else {
			count += 1
		}
	}
	if count == 0 {
		return defaultBit
	}
	if count > 0 {
		return '1'
	}
	return '0'
}

func filterByBit(input []string, idx int, filter rune) []string {
	filtered := []string{}
	for _, line := range input {
		if rune(line[idx]) == filter {
			filtered = append(filtered, line)
		}
	}
	return filtered
}

func filterByMostCommonBits(input []string) string {
	filtered := make([]string, len(input))
	copy(filtered, input)

	for idx := 0; idx < len(input[0]); idx++ {
		filterBy := getMostCommonBit(filtered, idx, '1')
		filtered = filterByBit(filtered, idx, filterBy)
		if len(filtered) == 1 {
			return filtered[0]
		}
	}

	panic("There were either none or more than 1 left at the end :(")
}

func filterByLeastCommonBits(input []string) string {
	filtered := make([]string, len(input))
	copy(filtered, input)

	for idx := 0; idx < len(input[0]); idx++ {
		filterBy := '1'
		if getMostCommonBit(filtered, idx, '1') == '1' {
			filterBy = '0'
		}
		filtered = filterByBit(filtered, idx, filterBy)
		if len(filtered) == 1 {
			return filtered[0]
		}
	}

	panic("There were either none or more than 1 left at the end :(")
}

func (sub *Submarine) ReadDiagnosticsReport(input []string) {
	sub.GammaRate = ""
	sub.EpsilonRate = ""

	for idx := 0; idx < len(input[0]); idx++ {
		if getMostCommonBit(input, idx, '1') == '1' {
			sub.GammaRate += "1"
			sub.EpsilonRate += "0"
		} else {
			sub.GammaRate += "0"
			sub.EpsilonRate += "1"
		}
	}

	sub.OxygenGeneratorRating = filterByMostCommonBits(input)
	sub.CO2ScrubberRating = filterByLeastCommonBits(input)
}

type SubmarineWithAim struct {
	*Submarine
	Aim int
}

func NewWithAim() *SubmarineWithAim {
	return &SubmarineWithAim{New(), 0}
}

func (sub *SubmarineWithAim) Run(cmd Command) {
	if cmd.Name == "up" {
		sub.Aim -= cmd.Value
	} else if cmd.Name == "down" {
		sub.Aim += cmd.Value
	} else if cmd.Name == "forward" {
		sub.Horizontal += cmd.Value
		sub.Depth += cmd.Value * sub.Aim
	} else {
		panic("Unknown command :(")
	}
}

type Command struct {
	Name  string
	Value int
}

func ParseCommand(input string) Command {
	split := strings.Fields(input)
	name, value := split[0], convert.ToInt(split[1])
	return Command{name, value}
}
