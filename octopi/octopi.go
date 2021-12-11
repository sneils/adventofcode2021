package octopi

import (
	"fmt"
	"log"

	"github.com/sneils/adventofcode2021/convert"
)

type Octopus struct {
	x, y, energy int
}

type OctoGrid struct {
	octopi                       [][]Octopus
	Age, Flashes, FlashesLastDay int
	lenX, lenY                   int
}

func Parse(inputs []string) *OctoGrid {
	grid := [][]Octopus{}
	for y, input := range inputs {
		row := []Octopus{}
		for x, char := range input {
			energy := convert.ToInt(string(char))
			octopus := Octopus{x, y, energy}
			row = append(row, octopus)
		}
		grid = append(grid, row)
	}
	return &OctoGrid{grid, 0, 0, 0, len(grid), len(grid[0])}
}

func printStringSlice(inputs []string) {
	for _, input := range inputs {
		log.Println(input)
	}
}

func contains(octopi []Octopus, find Octopus) bool {
	for _, octopus := range octopi {
		if octopus.x == find.x && octopus.y == find.y {
			return true
		}
	}
	return false
}

func (grid *OctoGrid) Pack() []string {
	packed := make([]string, grid.lenY)
	for y, row := range grid.octopi {
		for _, octopus := range row {
			packed[y] += fmt.Sprintf("%d", octopus.energy)
		}
	}
	return packed
}

func (grid *OctoGrid) energizeNeighbors(octopus Octopus) {
	x, y := octopus.x, octopus.y
	// left
	if x > 0 {
		grid.octopi[y][x-1].energy += 1
	}
	// top left
	if x > 0 && y > 0 {
		grid.octopi[y-1][x-1].energy += 1
	}
	// top
	if y > 0 {
		grid.octopi[y-1][x].energy += 1
	}
	// top right
	if x < grid.lenX-1 && y > 0 {
		grid.octopi[y-1][x+1].energy += 1
	}
	// right
	if x < grid.lenX-1 {
		grid.octopi[y][x+1].energy += 1
	}
	// bottom right
	if x < grid.lenX-1 && y < grid.lenY-1 {
		grid.octopi[y+1][x+1].energy += 1
	}
	// bottom
	if y < grid.lenY-1 {
		grid.octopi[y+1][x].energy += 1
	}
	// bottom left
	if y < grid.lenY-1 && x > 0 {
		grid.octopi[y+1][x-1].energy += 1
	}
}

func (grid *OctoGrid) DoAge() int {
	for x, row := range grid.octopi {
		for y := range row {
			grid.octopi[y][x].energy += 1
		}
	}
	flashed := []Octopus{}
	repeat := true
	for repeat {
		repeat = false
		for _, row := range grid.octopi {
			for _, octopus := range row {
				if contains(flashed, octopus) {
					continue
				}
				if octopus.energy <= 9 {
					continue
				}
				repeat = true
				flashed = append(flashed, octopus)
				grid.energizeNeighbors(octopus)
			}
		}
	}
	for _, octopus := range flashed {
		grid.octopi[octopus.y][octopus.x].energy = 0
	}
	grid.Age += 1
	grid.FlashesLastDay = len(flashed)
	grid.Flashes += grid.FlashesLastDay
	return grid.FlashesLastDay
}

func (grid *OctoGrid) ForwardToFirstFullFlash() {
	for grid.FlashesLastDay != 100 {
		grid.DoAge()
	}
}
