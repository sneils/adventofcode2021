package vents

import (
	"fmt"
	"strings"
)

const STRING_FORMAT = "%d,%d -> %d,%d"

type Vent struct {
	X1, X2 int
	Y1, Y2 int
}

func New(x1, y1, x2, y2 int) Vent {
	return Vent{X1: x1, Y1: y1, X2: x2, Y2: y2}
}

func ParseString(input string) Vent {
	var x1, y1, x2, y2 int
	reader := strings.NewReader(input)
	_, err := fmt.Fscanf(reader, STRING_FORMAT, &x1, &y1, &x2, &y2)
	if err != nil {
		panic("Invalid format :(")
	}
	return New(x1, y1, x2, y2)
}

func (vent Vent) String() string {
	return fmt.Sprintf(STRING_FORMAT, vent.X1, vent.Y1, vent.X2, vent.Y2)
}
