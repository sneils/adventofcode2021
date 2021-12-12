package passages

import (
	"fmt"
	"regexp"
	"strings"
)

type Tunnel struct {
	in, out string
}

func ParseTunnel(input string) Tunnel {
	inAndOut := strings.Split(input, "-")
	if inAndOut[0] == "end" || inAndOut[1] == "start" {
		return ParseTunnelReverse(input)
	}
	return Tunnel{in: inAndOut[0], out: inAndOut[1]}
}

func ParseTunnelReverse(input string) Tunnel {
	inAndOut := strings.Split(input, "-")
	return Tunnel{in: inAndOut[1], out: inAndOut[0]}
}

func (tunnel Tunnel) String() string {
	return fmt.Sprintf("%s-%s", tunnel.in, tunnel.out)
}

var BIG_CAVE_REGEXP = regexp.MustCompile(`[A-Z]`)

func (tunnel Tunnel) leadsToBigCave() bool {
	return BIG_CAVE_REGEXP.MatchString(tunnel.out)
}

type Passage []Tunnel

func ParsePassage(input string) Passage {
	parts := strings.Split(input, ",")
	passage := Passage{}
	for i := 0; i < len(parts)-1; i++ {
		tunnel := fmt.Sprintf("%s-%s", parts[i], parts[i+1])
		passage = append(passage, ParseTunnel(tunnel))
	}
	return passage
}

func (passage Passage) String() string {
	output := ""
	for _, tunnel := range passage {
		output += tunnel.in + ","
	}
	return output + passage[len(passage)-1].out
}

func (passage Passage) isVisitAllowed(tunnel Tunnel, allowOneRevisit bool) bool {
	if tunnel.leadsToBigCave() {
		return true
	}
	smallCavesVisited := map[string]int{}
	for _, visited := range passage {
		if visited.leadsToBigCave() {
			continue
		}
		smallCavesVisited[visited.out] += 1
	}
	if _, ok := smallCavesVisited[tunnel.out]; !ok {
		return true
	}
	if allowOneRevisit {
		for _, visits := range smallCavesVisited {
			if visits > 1 {
				return false
			}
		}
		return true
	}
	return false
}

type CaveSystem []Tunnel

func ParseCave(inputs []string) CaveSystem {
	cave := CaveSystem{}
	for _, input := range inputs {
		tunnel := ParseTunnel(input)
		cave = append(cave, tunnel)
		if tunnel.in == "start" || tunnel.out == "end" {
			continue
		}
		cave = append(cave, ParseTunnelReverse(input))
	}
	return cave
}

func (cave CaveSystem) getNextTunnels(start Tunnel) []Tunnel {
	next := []Tunnel{}
	for _, tunnel := range cave {
		if tunnel.in == start.out {
			next = append(next, tunnel)
		}
	}
	return next
}

func (cave CaveSystem) findPassages(passage Passage, allowSingleRevisit bool) []Passage {
	head := passage[len(passage)-1]
	body := passage[:len(passage)-1]
	if head.out == "end" {
		return []Passage{passage}
	}
	if !body.isVisitAllowed(head, allowSingleRevisit) {
		return []Passage{}
	}
	passages := []Passage{}
	next := cave.getNextTunnels(head)
	for _, tunnel := range next {
		nextPassage := make(Passage, len(passage))
		copy(nextPassage, passage)
		nextPassage = append(nextPassage, tunnel)
		passages = append(passages, cave.findPassages(nextPassage, allowSingleRevisit)...)
	}
	return passages
}

func (cave CaveSystem) CountPassagesOut(allowSingleRevisit bool) int {
	count := 0
	for _, tunnel := range cave {
		if tunnel.in != "start" {
			continue
		}
		count += len(cave.findPassages(Passage{tunnel}, allowSingleRevisit))
	}
	return count
}
