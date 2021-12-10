package ints

import "testing"

func TestSum(t *testing.T) {
	expected := 144
	sum := GetSum([]int{1, 2, 118, 3, 4, 5, 11})
	if expected != sum {
		t.Errorf("Expected sum to be %d, found %d.", expected, sum)
	}
}

func TestGetTriangleNumber(t *testing.T) {
	expected := map[int]int{
		1: 1,
		2: 3,
		3: 6,
		4: 10,
		5: 15,
		6: 21,
	}
	for key, val := range expected {
		num := GetTriangleNumber(key)
		if val != num {
			t.Errorf("Expected triangle number for %d to be %d, found %d.", key, val, num)
		}
	}
}

func TestSumTriangleNumbers(t *testing.T) {
	expected := 20
	num := SumTriangleNumbers(1, []int{1, 2, 3, 4, 5})
	if expected != num {
		t.Errorf("Expected sum to be %d, found %d.", expected, num)
	}
}

func TestGetMedianOdd(t *testing.T) {
	expected := 2
	ints := []int{1, 2, 3}
	median := GetMedian(ints)
	if expected != median {
		t.Errorf("Expected median to be %d, found %d.", expected, median)
	}
}

func TestGetMedianEven1223(t *testing.T) {
	expected := 2
	ints := []int{1, 2, 2, 3}
	median := GetMedian(ints)
	if expected != median {
		t.Errorf("Expected median to be %d, found %d.", expected, median)
	}
}

func TestGetMedianEven1155(t *testing.T) {
	expected := 3
	ints := []int{1, 1, 5, 5}
	median := GetMedian(ints)
	if expected != median {
		t.Errorf("Expected median to be %d, found %d.", expected, median)
	}
}

func TestGetMean(t *testing.T) {
	expected := 4
	ints := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	mean := GetMean(ints)
	if expected != mean {
		t.Errorf("Expected mean to be %d, found %d.", expected, mean)
	}
}
