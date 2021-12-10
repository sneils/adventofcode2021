package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/sneils/adventofcode2021/days"
)

const LAST = 10

func main() {
	for day := 1; day <= LAST; day++ {
		returns := reflect.ValueOf(days.New(day)).MethodByName(fmt.Sprintf("Day%d", day)).Call([]reflect.Value{})
		log.Printf("day: %2d, part1: %15d, part2: %15d\n", day, returns[0].Int(), returns[1].Int())
	}
}
