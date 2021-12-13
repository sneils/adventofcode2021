package origami

import (
	"fmt"
	"log"
	"strings"

	"github.com/sneils/adventofcode2021/convert"
)

type Paper map[string]bool

func Parse(inputs []string) Paper {
	paper := Paper{}
	for _, input := range inputs {
		paper[input] = true
	}
	return paper
}

func (paper Paper) CountDots() int {
	return len(paper)
}

func getXY(key string) (int, int) {
	split := strings.Split(key, ",")
	ints := convert.ToInts(split)
	return ints[0], ints[1]
}

func getKey(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func (paper Paper) Set(x, y int) {
	paper[getKey(x, y)] = true
}

func ParseFold(input string) (string, int) {
	fields := strings.Fields(input)
	split := strings.Split(fields[2], "=")
	return split[0], convert.ToInt(split[1])
}

func (paper Paper) Fold(input string) {
	dir, pos := ParseFold(input)
	for dot := range paper {
		x, y := getXY(dot)
		if dir == "x" && pos < x {
			paper.Set(pos*2-x, y)
			delete(paper, dot)
		}
		if dir == "y" && pos < y {
			paper.Set(x, pos*2-y)
			delete(paper, dot)
		}
	}
}

func (paper Paper) Print() {
	maxX, maxY := 0, 0
	for dot := range paper {
		x, y := getXY(dot)
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}
	for y := 0; y <= maxY; y++ {
		row := ""
		for x := 0; x <= maxX; x++ {
			if _, ok := paper[getKey(x, y)]; ok {
				row += "#"
			} else {
				row += "."
			}
		}
		log.Println(row)
	}
}
