package lavatubes

import (
	"sort"

	"github.com/sneils/adventofcode2021/convert"
)

type LavaTube struct {
	x, y, z int
}

type Region struct {
	tubes []LavaTube
	x, y  int
}

func Parse(inputs []string) Region {
	region := Region{}
	region.x = len(inputs[0])
	region.y = len(inputs)
	for y, row := range inputs {
		for x, col := range row {
			z := convert.ToInt(string(col))
			tube := LavaTube{x, y, z}
			region.tubes = append(region.tubes, tube)
		}
	}
	return region
}

func (region Region) Get(x, y int) LavaTube {
	for _, tube := range region.tubes {
		if tube.x == x && tube.y == y {
			return tube
		}
	}
	panic("LavaTube not found :x")
}

func (region Region) getNeighbors(tube LavaTube) []LavaTube {
	neighbors := []LavaTube{}
	x, y := tube.x, tube.y
	if y > 0 {
		neighbor := region.Get(x, y-1)
		neighbors = append(neighbors, neighbor)
	}
	if x > 0 {
		neighbor := region.Get(x-1, y)
		neighbors = append(neighbors, neighbor)
	}
	if x < region.x-1 {
		neighbor := region.Get(x+1, y)
		neighbors = append(neighbors, neighbor)
	}
	if y < region.y-1 {
		neighbor := region.Get(x, y+1)
		neighbors = append(neighbors, neighbor)
	}
	return neighbors
}

func (region Region) getLowpoints() []LavaTube {
	lowpoints := []LavaTube{}

	for _, tube := range region.tubes {
		isLowpoint := true
		for _, neighbor := range region.getNeighbors(tube) {
			if neighbor.z <= tube.z {
				isLowpoint = false
				break
			}
		}
		if isLowpoint {
			lowpoints = append(lowpoints, tube)
		}
	}
	return lowpoints
}

func (region Region) CalculateRisk() int {
	count := 0
	for _, lowpoint := range region.getLowpoints() {
		count += lowpoint.z + 1
	}
	return count
}

func contains(tubes []LavaTube, find LavaTube) bool {
	for _, tube := range tubes {
		if find.x == tube.x && find.y == tube.y {
			return true
		}
	}
	return false
}

func (region Region) getBasin(tube LavaTube, basin []LavaTube) []LavaTube {
	if tube.z == 9 {
		return basin
	}

	if contains(basin, tube) {
		return basin
	}

	basin = append(basin, tube)

	for _, neighbor := range region.getNeighbors(tube) {
		basin = region.getBasin(neighbor, basin)
	}

	return basin
}

func (region Region) getBasins() [][]LavaTube {
	basins := [][]LavaTube{}
	for _, lowpoint := range region.getLowpoints() {
		basin := region.getBasin(lowpoint, []LavaTube{})
		basins = append(basins, basin)
	}
	return basins
}

func (region Region) GetBiggestThreeBasinSizes() []int {
	sizes := []int{}
	for _, basin := range region.getBasins() {
		sizes = append(sizes, len(basin))
	}
	sort.Ints(sizes)
	return sizes[len(sizes)-3:]
}
