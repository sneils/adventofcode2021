package main

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/sneils/adventofcode2021/days"
)

func main() {
	for day := 1; day <= 13; day++ {
		now := time.Now()
		returns := reflect.ValueOf(days.New(day)).MethodByName(fmt.Sprintf("Day%d", day)).Call([]reflect.Value{})
		log.Printf("day: %2d, part1: %15d, part2: %15d, time: %v\n", day, returns[0].Int(), returns[1].Int(), time.Since(now))
	}
}
