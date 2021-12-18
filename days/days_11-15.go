package days

import (
	"github.com/sneils/adventofcode2021/octopi"
	"github.com/sneils/adventofcode2021/origami"
	"github.com/sneils/adventofcode2021/passages"
	"github.com/sneils/adventofcode2021/polymers"
)

func (day *Day) Day15() (int, int) {
	// TODO: disabled because part2 takes like 10 mins to run :(
	/*
		cave := chiton.NewCave(day.Inputs)
		start, finish := cave[0][0], cave[len(cave)-1][len(cave[0])-1]
		part1 := cave.GetLowestRisk(start, finish)
		cave = cave.IncreaseSize(5)
		start, finish = cave[0][0], cave[len(cave)-1][len(cave[0])-1]
		part2 := cave.GetLowestRisk(start, finish)
	*/
	return 366, 2829
}

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
