package polymers

import (
	"fmt"
)

type Template struct {
	first, last byte
	pairs       map[string]int
}

func NewTemplate(input string) *Template {
	tmpl := Template{
		first: input[0],
		last:  input[len(input)-1],
		pairs: map[string]int{},
	}
	for i := 0; i < len(input)-1; i++ {
		tmpl.pairs[input[i:i+2]] += 1
	}
	return &tmpl
}

type RuleSet map[string]string

func NewRuleSet(inputs []string) RuleSet {
	rules := RuleSet{}
	for _, input := range inputs {
		rules[input[:2]] = string(input[6])
	}
	return rules
}

func (tmpl *Template) ApplyRules(rules RuleSet) {
	add := map[string]int{}
	for find, insert := range rules {
		if cnt, ok := tmpl.pairs[find]; ok {
			delete(tmpl.pairs, find)
			add[fmt.Sprintf("%c%s", find[0], insert)] += cnt
			add[fmt.Sprintf("%s%c", insert, find[1])] += cnt
		}
	}
	for key, val := range add {
		tmpl.pairs[key] = val
	}
}

func (tmpl *Template) countElements() map[byte]int {
	elems := map[byte]int{}
	for key, val := range tmpl.pairs {
		elems[key[0]] += val
		elems[key[1]] += val
	}
	for key := range elems {
		elems[key] /= 2
	}
	elems[tmpl.first] += 1
	elems[tmpl.last] += 1
	return elems
}

func (tmpl *Template) GetScore() int {
	min, max := -1, -1
	for _, val := range tmpl.countElements() {
		if min < 0 || min > val {
			min = val
		}
		if max < 0 || max < val {
			max = val
		}
	}
	return max - min
}
