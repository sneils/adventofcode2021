package lanternfish

import (
	"strings"

	"github.com/sneils/adventofcode2021/convert"
)

const (
	NEWBORN_AGE = 8
	RESET_AGE   = 6
)

type FishSchool struct {
	members map[int]int
	Age     int
}

func ParseString(inputs string) *FishSchool {
	school := &FishSchool{members: map[int]int{}, Age: 0}
	for _, input := range strings.Split(inputs, ",") {
		i := convert.ToInt(input)
		school.members[i] += 1
	}
	return school
}

func (school *FishSchool) DoAge() {
	older := map[int]int{}
	for key, val := range school.members {
		if key == 0 {
			older[NEWBORN_AGE] += val
			older[RESET_AGE] += val
		} else {
			older[key-1] += val
		}
	}
	school.Age++
	school.members = older
}

func (school *FishSchool) Get(i int) int {
	return school.members[i]
}

func (school *FishSchool) Count() int {
	count := 0
	for _, val := range school.members {
		count += val
	}
	return count
}
