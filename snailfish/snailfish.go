package snailfish

import (
	"fmt"
	"strings"

	"github.com/sneils/adventofcode2021/convert"
)

type Number struct {
	input  string
	a, b   int
	an, bn *Number
}

func NewNumber(input string) *Number {
	// strip outer []
	inner := input[1 : len(input)-1]
	a, an := parse(inner)
	if an == nil {
		inner = inner[len(fmt.Sprintf("%d", a))+1:]
	} else {
		inner = inner[len(an.input)+1:]
	}
	b, bn := parse(inner)
	return &Number{a: a, b: b, an: an, bn: bn, input: input}
}

func (n Number) String() string {
	return n.input
}

func (n *Number) Reduce() *Number {
	next := n.Explode()
	if next.String() != n.String() {
		return next.Reduce()
	}
	next = n.Split()
	if next.String() != n.String() {
		return next.Reduce()
	}
	return next
}

func (n *Number) Explode() *Number {
	str := n.String()
	stack := 0
	for i, check := range str {
		if check == '[' {
			stack++
		}
		if check == ']' {
			stack--
		}
		if stack == 5 {
			last := getLastNumber(str[:i])
			this := next(str[i:])
			expl := NewNumber(this)
			next := getNextNumber(str[i+len(this):])

			a := str[:i]
			skipA := strings.LastIndex(a, last)
			if last != "" {
				an := convert.ToInt(last) + expl.a
				a = fmt.Sprintf("%s%d%s", a[:skipA], an, a[len(last)+skipA:])
			}
			b := str[i+len(this):]
			skipB := strings.Index(b, next)
			if next != "" {
				bn := convert.ToInt(next) + expl.b
				b = fmt.Sprintf("%s%d%s", b[:skipB], bn, b[skipB+len(next):])
			}
			return NewNumber(a + "0" + b)
		}
	}
	return n
}

func (n *Number) Split() *Number {
	done := ""
	todo := n.String()
	for len(todo) > 0 {
		next := getNextNumber(todo)
		if len(next) < 1 {
			break
		}
		skip := strings.Index(todo, next)
		done += todo[:skip]
		todo = todo[skip+len(next):]
		if len(next) == 1 {
			done += next
			continue
		}
		i := convert.ToInt(next)
		x := Split(i).String()
		return NewNumber(done + x + todo)
	}
	return n
}

func (n *Number) Add(a *Number) *Number {
	return NewNumber(fmt.Sprintf("[%s,%s]", n, a)).Reduce()
}

func (n *Number) GetMagnitude() int {
	sum := 0
	if n.a > -1 {
		sum += 3 * n.a
	} else {
		sum += 3 * n.an.GetMagnitude()
	}
	if n.b > -1 {
		sum += 2 * n.b
	} else {
		sum += 2 * n.bn.GetMagnitude()
	}
	return sum
}

type Numbers []*Number

func (n Numbers) Sum() *Number {
	a := n[0]
	for _, b := range n[1:] {
		a = a.Add(b)
	}
	return a
}

func (n Numbers) GetBestMagnitude() int {
	best := 0
	for ai, a := range n {
		for bi, b := range n {
			if ai == bi {
				continue
			}
			ab := a.Add(b)
			abm := ab.GetMagnitude()
			if abm > best {
				best = abm
			}
			ba := b.Add(a)
			bam := ba.GetMagnitude()
			if bam > best {
				best = bam
			}
		}
	}
	return best
}
