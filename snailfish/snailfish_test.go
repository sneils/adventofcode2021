package snailfish

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []string{
		"[1,2]",
		"[[1,2],3]",
		"[9,[8,7]]",
		"[[1,9],[8,5]]",
		"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]",
		"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]",
		"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
		"[11,22]",
	}
	for _, input := range tests {
		n := NewNumber(input)
		if input != n.String() {
			t.Errorf("Expected to look like %s, but looks like %s.", input, n)
		}
	}
}

func TestSplit2(t *testing.T) {
	tests := map[string]string{
		"[[[[0,7],4],[15,[0,13]]],[1,1]]":    "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
		"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]": "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
	}
	for input, expected := range tests {
		n := NewNumber(input)
		exploded := n.Split()
		if exploded.String() != expected {
			t.Errorf("Expected split number to look like %s, found %s.", expected, exploded)
		}
	}
}

func TestExplode(t *testing.T) {
	tests := map[string]string{
		"[[[[[9,8],1],2],3],4]":                 "[[[[0,9],2],3],4]", // (the 9 has no regular number to its left, so it is not added to any regular number).
		"[7,[6,[5,[4,[3,2]]]]]":                 "[7,[6,[5,[7,0]]]]", // (the 2 has no regular number to its right, and so it is not added to any regular number).
		"[[6,[5,[4,[3,2]]]],1]":                 "[[6,[5,[7,0]]],3]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]": "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", // (the pair [3,2] is unaffected because the pair [7,3] is further to the left; [3,2] would explode on the next action).
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]":     "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		"[[[[0,7],4],[7,[[8,4],9]]],[1,1]]":     "[[[[0,7],4],[15,[0,13]]],[1,1]]",
		"[[[[0,7],4],[15,[0,13]]],[1,1]]":       "[[[[0,7],4],[15,[0,13]]],[1,1]]",
	}
	for input, expected := range tests {
		n := NewNumber(input)
		actual := n.Explode()
		if expected != actual.String() {
			t.Errorf("Expected exploded number to look like %s, found %s.", expected, actual)
		}
	}
}

func TestAdd(t *testing.T) {
	a, b := NewNumber("[[[[4,3],4],4],[7,[[8,4],9]]]"), NewNumber("[1,1]")
	expected := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
	actual := a.Add(b)
	if expected != actual.String() {
		t.Errorf("Expected result to be %s, found %s.", expected, actual)
	}
}

func TestSum(t *testing.T) {
	adds := []string{
		"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
		"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
		"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
		"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
		"[7,[5,[[3,8],[1,4]]]]",
		"[[2,[2,2]],[8,[8,1]]]",
		"[2,9]",
		"[1,[[[9,3],9],[[9,0],[0,7]]]]",
		"[[[5,[7,4]],7],1]",
		"[[[[4,2],2],6],[8,7]]",
	}
	expected := "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"
	numbers := Numbers{}
	for _, input := range adds {
		n := NewNumber(input)
		numbers = append(numbers, n)
	}
	actual := numbers.Sum()
	if expected != actual.String() {
		t.Errorf("Expected result to be %s, found %s.", expected, actual)
	}
}

func TestReduce(t *testing.T) {
	input := "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"
	expected := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
	actual := NewNumber(input).Reduce()
	if expected != actual.String() {
		t.Errorf("Expected reduction to look like %s, found %s.", expected, actual)
	}
}

func TestGetMagnitude(t *testing.T) {
	tests := map[string]int{
		"[[1,2],[[3,4],5]]":                                             143,
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]":                             1384,
		"[[[[1,1],[2,2]],[3,3]],[4,4]]":                                 445,
		"[[[[3,0],[5,3]],[4,4]],[5,5]]":                                 791,
		"[[[[5,0],[7,4]],[5,5]],[6,6]]":                                 1137,
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]":         3488,
		"[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]": 4140,
	}
	for input, expected := range tests {
		n := NewNumber(input)
		actual := n.GetMagnitude()
		if expected != actual {
			t.Errorf("Expected magnitude to be %d, found %d.", expected, actual)
		}
	}
}

func TestGetBestMagnitude(t *testing.T) {
	adds := []string{
		"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
		"[[[5,[2,8]],4],[5,[[9,9],0]]]",
		"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
		"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
		"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
		"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
		"[[[[5,4],[7,7]],8],[[8,3],8]]",
		"[[9,3],[[9,9],[6,[4,9]]]]",
		"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
		"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
	}
	expected := 3993
	numbers := Numbers{}
	for _, input := range adds {
		n := NewNumber(input)
		numbers = append(numbers, n)
	}
	actual := numbers.GetBestMagnitude()
	if expected != actual {
		t.Errorf("Expected result to be %d, found %d.", expected, actual)
	}
}
