package ventmap

import (
	"fmt"

	"github.com/sneils/adventofcode2021/vents"
)

type VentMap struct {
	grid          map[string]int
	withDiagonals bool
}

func New(withDiagonals bool) *VentMap {
	return &VentMap{grid: map[string]int{}, withDiagonals: withDiagonals}
}

func (vm *VentMap) Get(x, y int) int {
	return vm.grid[fmt.Sprintf("%d,%d", x, y)]
}

func (vm *VentMap) Add(vent vents.Vent) {
	xd, xn := vent.X2-vent.X1, 0
	if xd > 0 {
		xn += 1
	} else if xd < 0 {
		xn -= 1
		xd *= -1
	}

	yd, yn := vent.Y2-vent.Y1, 0
	if yd > 0 {
		yn += 1
	} else if yd < 0 {
		yn -= 1
		yd *= -1
	}

	if !vm.withDiagonals && xd > 0 && yd > 0 {
		return
	}

	for i := 0; i <= xd || i <= yd; i++ {
		vm.grid[fmt.Sprintf("%d,%d", vent.X1+i*xn, vent.Y1+i*yn)] += 1
	}
}

func (vm *VentMap) GetOverlaps() int {
	score := 0
	for _, val := range vm.grid {
		if val > 1 {
			score += 1
		}
	}
	return score
}
