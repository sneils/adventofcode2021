package ints

import (
	"sort"
)

func GetMedian(ints []int) int {
	sort.Ints(ints)
	l := len(ints)
	if l%2 > 0 {
		return ints[l/2]
	}
	return GetMean([]int{ints[l/2], ints[l/2-1]})
}

func GetSum(ints []int) int {
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return sum
}

func GetMean(ints []int) int {
	return GetSum(ints) / len(ints)
}

func GetTriangleNumber(n int) int {
	sum := 0
	if n < 0 {
		n *= -1
	}
	for i := 0; i < n; i++ {
		sum += i + 1
	}
	return sum
}

func SumTriangleNumbers(n int, from []int) int {
	sum := 0
	for _, i := range from {
		sum += GetTriangleNumber(n - i)
	}
	return sum
}
