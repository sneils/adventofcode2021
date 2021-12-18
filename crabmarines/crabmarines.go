package crabmarines

import (
	"strings"

	"github.com/sneils/adventofcode2021/convert"
	"github.com/sneils/adventofcode2021/ints"
)

type CrabSubmarines struct {
	positions []int
}

func NewCrabmarines(inputs string) *CrabSubmarines {
	crabs := &CrabSubmarines{}
	for _, input := range strings.Split(inputs, ",") {
		i := convert.ToInt(input)
		crabs.positions = append(crabs.positions, i)
	}
	return crabs
}

func (crabs *CrabSubmarines) Count() int {
	return len(crabs.positions)
}

func (crabs *CrabSubmarines) AlignConstant() (int, int) {
	fuel, median := 0, ints.GetMedian(crabs.positions)
	for _, pos := range crabs.positions {
		diff := median - pos
		if diff < 0 {
			diff *= -1
		}
		fuel += diff
	}
	return median, fuel
}

func (crabs *CrabSubmarines) AlignIncreasing() (int, int) {
	mean := ints.GetMean(crabs.positions)
	fuel := ints.SumTriangleNumbers(mean, crabs.positions)

	// as our mean function always rounds down,
	// the next position might actually be the ideal one
	meanP1 := mean + 1
	fuelP1 := ints.SumTriangleNumbers(meanP1, crabs.positions)

	if fuelP1 < fuel {
		return meanP1, fuelP1
	}
	return mean, fuel
}
