package convert

import "strconv"

func ToIntBase(input string, base int) int {
	i, err := strconv.ParseInt(input, base, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func ToIntsBase(inputs []string, base int) []int {
	ints := []int{}
	for _, input := range inputs {
		i := ToIntBase(input, base)
		ints = append(ints, i)
	}
	return ints
}

func ToInt(input string) int {
	return ToIntBase(input, 10)
}

func ToInts(inputs []string) []int {
	return ToIntsBase(inputs, 10)
}

func FromBinary(input string) int {
	return ToIntBase(input, 2)
}

func FromHex(input string) int {
	return ToIntBase(input, 16)
}
