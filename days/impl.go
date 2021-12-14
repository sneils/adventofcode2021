package days

import (
	"strings"

	"github.com/sneils/adventofcode2021/bingo"
	"github.com/sneils/adventofcode2021/convert"
	"github.com/sneils/adventofcode2021/crabsubmarines"
	"github.com/sneils/adventofcode2021/display"
	"github.com/sneils/adventofcode2021/lanternfish"
	"github.com/sneils/adventofcode2021/lavatubes"
	"github.com/sneils/adventofcode2021/navigation"
	"github.com/sneils/adventofcode2021/octopi"
	"github.com/sneils/adventofcode2021/origami"
	"github.com/sneils/adventofcode2021/passages"
	"github.com/sneils/adventofcode2021/polymers"
	"github.com/sneils/adventofcode2021/submarines"
	"github.com/sneils/adventofcode2021/vents"
	"github.com/sneils/adventofcode2021/vents/ventmap"
)

func (day *Day) Day14() (int, int) {
	groups := day.Groups()
	tmpl := polymers.NewTemplate(groups[0][0])
	rules := polymers.NewRuleSet(groups[1])
	part1 := 0
	for i := 1; i <= 40; i++ {
		tmpl.ApplyRules(rules)
		if i == 10 {
			part1 = tmpl.GetScore()
		}
	}
	part2 := tmpl.GetScore()
	return part1, part2
}

func (day *Day) Day13() (int, int) {
	groups := day.Groups()
	paper := origami.Parse(groups[0])
	for _, fold := range groups[1][:1] {
		paper.Fold(fold)
	}
	part1 := paper.CountDots()
	for _, fold := range groups[1][1:] {
		paper.Fold(fold)
	}
	part2 := paper.CountDots()
	paper.Print()
	return part1, part2
}

func (day *Day) Day12() (int, int) {
	cave := passages.NewCaveSystem(day.Inputs)
	return cave.CountPassagesOut(false), cave.CountPassagesOut(true)
}

func (day *Day) Day11() (int, int) {
	grid := octopi.Parse(day.Inputs)
	for grid.Age < 100 {
		grid.DoAge()
	}
	part1 := grid.Flashes
	grid.ForwardToFirstFullFlash()
	part2 := grid.Age
	return part1, part2
}

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
	crabs := crabsubmarines.ParseString(day.Inputs[0])
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

func (day *Day) Day5() (int, int) {
	part1, part2 := ventmap.New(false), ventmap.New(true)
	for _, input := range day.Inputs {
		vent := vents.ParseString(input)
		part1.Add(vent)
		part2.Add(vent)
	}
	return part1.GetOverlaps(), part2.GetOverlaps()
}

func (day *Day) Day4() (int, int) {
	groups := day.Groups()
	picks := strings.Split(groups[0][0], ",")
	boards := []bingo.Board{}
	for _, group := range groups[1:] {
		board := bingo.New(group)
		boards = append(boards, board)
	}
	return Day4Part1(boards, picks), Day4Part2(boards, picks)
}

func Day4Part1(boards []bingo.Board, picks []string) int {
	for _, pick := range picks {
		for i := 0; i < len(boards); i++ {
			if boards[i].Mark(pick) && boards[i].HasWon() {
				return boards[i].GetScore() * convert.ToInt(pick)
			}
		}
	}
	panic("Nobody won??")
}

func Day4Part2(boards []bingo.Board, picks []string) int {
	for _, pick := range picks {
		losers := []bingo.Board{}
		for i := 0; i < len(boards); i++ {
			if boards[i].Mark(pick) && boards[i].HasWon() {
				continue
			}
			losers = append(losers, boards[i])
		}
		if len(losers) == 0 {
			return boards[0].GetScore() * convert.ToInt(pick)
		}
		boards = losers
	}
	panic("What the f is going on?")
}

func (day *Day) Day3() (int, int) {
	sub := submarines.New()
	sub.ReadDiagnosticsReport(day.Inputs)
	return sub.GetPowerConsumption(), sub.GetLifeSupportRating()
}

func (day *Day) Day2() (int, int) {
	sub1, sub2 := submarines.New(), submarines.NewWithAim()
	for _, line := range day.Inputs {
		cmd := submarines.ParseCommand(line)
		sub1.Run(cmd)
		sub2.Run(cmd)
	}
	return sub1.GetPosition(), sub2.GetPosition()
}

func (day *Day) Day1() (int, int) {
	ints := convert.ToInts(day.Inputs)

	part1, part2 := 0, 0
	for idx := 1; idx < len(ints); idx++ {
		if ints[idx-1] < ints[idx] {
			part1++
		}
		if idx > len(ints)-3 {
			continue
		}
		if ints[idx-1] < ints[idx+2] {
			part2++
		}
	}
	return part1, part2
}
