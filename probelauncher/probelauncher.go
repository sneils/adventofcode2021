package probelauncher

import (
	"fmt"
	"strings"
)

type Probe struct {
	x, y int
}

func NewProbe(x, y int) *Probe {
	return &Probe{x: x, y: y}
}

func (p *Probe) LandsIn(t *Target) (bool, int) {
	for i, x, y, maxY := 0, 0, 0, 0; ; i += 1 {
		if p.x-i > 0 {
			x += p.x - i
		}
		y += p.y - i
		if y > maxY {
			maxY = y
		}
		if t.contains(x, y) {
			return true, maxY
		}
		if x > t.x2 || y < t.y1 {
			return false, maxY
		}
	}
}

const TARGET_FORMAT = "target area: x=%d..%d, y=%d..%d"

type Target struct {
	x1, x2, y1, y2 int
}

func NewTarget(input string) *Target {
	var x1, y1, x2, y2 int
	reader := strings.NewReader(input)
	_, err := fmt.Fscanf(reader, TARGET_FORMAT, &x1, &x2, &y1, &y2)
	if err != nil {
		panic("Invalid format :(")
	}
	return &Target{x1: x1, x2: x2, y1: y1, y2: y2}
}

func (t *Target) contains(x, y int) bool {
	return x >= t.x1 && x <= t.x2 && y >= t.y1 && y <= t.y2
}

func (t *Target) String() string {
	return fmt.Sprintf(TARGET_FORMAT, t.x1, t.x2, t.y1, t.y2)
}

func (t *Target) GetHits() TargetHits {
	hits := TargetHits{}
	minX, maxX := t.getMinX(), t.x2
	minY, maxY := t.y1, t.y1*-1
	for x := minX; x <= maxX; x += 1 {
		for y := minY; y <= maxY; y += 1 {
			probe := NewProbe(x, y)
			if ok, maxY := probe.LandsIn(t); ok {
				hit := TargetHit{
					probe: probe,
					maxY:  maxY,
				}
				hits = append(hits, hit)
			}
		}
	}
	return hits
}

func (t *Target) getMinX() int {
	for x := 1; x <= t.x2; x += 1 {
		sum, cur := 0, x
		for cur > 0 {
			sum += cur
			cur -= 1
		}
		if sum >= t.x1 {
			return x
		}
	}
	return 1
}

type TargetHit struct {
	probe *Probe
	maxY  int
}

type TargetHits []TargetHit

func (t TargetHits) GetMaxY() int {
	maxY := 0
	for _, hit := range t {
		if hit.maxY > maxY {
			maxY = hit.maxY
		}
	}
	return maxY
}
