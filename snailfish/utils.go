package snailfish

import (
	"fmt"
	"strings"
)

func format(i int, n *Number) string {
	if i >= 0 {
		return fmt.Sprintf("%d", i)
	}
	return n.String()
}

func next(input string) string {
	stack := 0
	str := ""
	for _, check := range input {
		if check == '[' {
			stack++
		}
		if check == ']' {
			stack--
		}
		str += string(check)
		if stack == 0 {
			break
		}
	}
	return str
}

func Split(n int) *Number {
	if n < 10 {
		panic("n < 10")
	}
	a := n / 2
	b := a
	if n%2 > 0 {
		b += 1
	}
	str := fmt.Sprintf("[%d,%d]", a, b)
	return NewNumber(str)
}

func parse(input string) (int, *Number) {
	if input[0] == '[' {
		return -1, NewNumber(next(input))
	}
	n := 0
	reader := strings.NewReader(input)
	_, err := fmt.Fscanf(reader, "%d", &n)
	if err == nil {
		return n, nil
	}
	panic(err)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getNextNumber(input string) string {
	str := ""
	started := false
	for i := 0; i < len(input); i += 1 {
		if input[i] == '[' || input[i] == ']' || input[i] == ',' {
			if started {
				break
			}
			continue
		}
		started = true
		str += string(input[i])
	}
	return str
}

func getLastNumber(input string) string {
	reversed := reverse(input)
	str := getNextNumber(reversed)
	return reverse(str)
}
