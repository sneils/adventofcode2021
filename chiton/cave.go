package chiton

import (
	"errors"
	"fmt"

	"github.com/sneils/adventofcode2021/convert"
)

type Cave [][]Node

func NewCave(inputs []string) Cave {
	cave := Cave{}
	for y, input := range inputs {
		list := []Node{}
		for x, field := range input {
			risk := convert.ToInt(string(field))
			node := Node{
				x:      x,
				y:      y,
				risk:   risk,
				parent: nil,
			}
			list = append(list, node)
		}
		cave = append(cave, list)
	}
	return cave
}

func (cave Cave) getTotalRisk() int {
	risk := 0
	for _, row := range cave {
		for _, node := range row {
			risk += node.risk
		}
	}
	return risk
}

func (cave Cave) IncreaseSize(n int) Cave {
	inputs := []string{}
	maxX, maxY := len(cave[0]), len(cave)
	for y := 0; y < maxY*n; y++ {
		input := ""
		for x := 0; x < maxX*n; x++ {
			risk := cave[y%maxY][x%maxX].risk + y/maxY + x/maxX
			if risk > 9 {
				risk -= 9
			}
			input += fmt.Sprintf("%d", risk)
		}
		inputs = append(inputs, input)
	}
	return NewCave(inputs)
}

func (cave Cave) get(x, y int) (Node, error) {
	if x < 0 || y < 0 || x > len(cave[0])-1 || y > len(cave)-1 {
		return Node{}, errors.New("out of range :(")
	}
	return cave[y][x], nil
}

func (cave Cave) getNeighbors(node Node) NodeList {
	list := NodeList{}
	if neighbor, err := cave.get(node.x, node.y-1); err == nil {
		list = append(list, neighbor)
	}
	if neighbor, err := cave.get(node.x, node.y+1); err == nil {
		list = append(list, neighbor)
	}
	if neighbor, err := cave.get(node.x-1, node.y); err == nil {
		list = append(list, neighbor)
	}
	if neighbor, err := cave.get(node.x+1, node.y); err == nil {
		list = append(list, neighbor)
	}
	for i := 0; i < len(list); i++ {
		list[i].parent = &node
	}
	return list
}

func (cave Cave) GetLowestRisk(start, finish Node) int {
	todo := NodeList{start}
	done := NodeList{}
	for len(todo) > 0 {
		node := todo.getLowestScore(finish)
		if node.x == finish.x && node.y == finish.y {
			return node.getRisk()
		}
		todo = todo.except(node)
		done = append(done, node)
		for _, neighbor := range cave.getNeighbors(node) {
			if done.contains(neighbor) {
				continue
			}
			if todo.contains(neighbor) {
				continue
			}
			todo = append(todo, neighbor)
		}
	}
	panic("No path found :*(")
}
