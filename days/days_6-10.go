package days

import (
	"github.com/sneils/adventofcode2021/crabmarines"
	"github.com/sneils/adventofcode2021/display"
	"github.com/sneils/adventofcode2021/lanternfish"
	"github.com/sneils/adventofcode2021/lavatubes"
	"github.com/sneils/adventofcode2021/navigation"
)

func (day *Day) Day10() (int, int) {
	return navigation.Analyse(day.Inputs)
}

func (day *Day) Day9() (int, int) {
	tubes := lavatubes.Parse(day.Inputs)
	basins := tubes.GetBiggestThreeBasinSizes()
	area := basins[0]
	for _, size := range basins[1:] {
		area *= size
	}
	return tubes.CalculateRisk(), area
}

func (day *Day) Day8() (int, int) {
	part1, part2 := 0, 0
	for _, input := range day.Inputs {
		d := display.ParseString(input)
		part1 += d.CountUniqueOutputForms()
		part2 += d.SumOutputs()
	}
	return part1, part2
}

func (day *Day) Day7() (int, int) {
	crabs := crabmarines.NewCrabmarines(day.Inputs[0])
	_, part1 := crabs.AlignConstant()
	_, part2 := crabs.AlignIncreasing()
	return part1, part2
}

func (day *Day) Day6() (int, int) {
	school := lanternfish.ParseString(day.Inputs[0])
	part1 := 0
	for {
		school.DoAge()
		if school.Age == 80 {
			part1 = school.Count()
		}
		if school.Age == 256 {
			return part1, school.Count()
		}
	}
}
