package chiton

type NodeList []Node

func (list NodeList) except(rm Node) NodeList {
	new := NodeList{}
	for _, node := range list {
		if node.x == rm.x && node.y == rm.y {
			continue
		}
		new = append(new, node)
	}
	return new
}

func (list NodeList) contains(find Node) bool {
	for _, node := range list {
		if node.x == find.x && node.y == find.y {
			return true
		}
	}
	return false
}

func (list NodeList) getLowestScore(finish Node) Node {
	min, minScore := list[0], list[0].getScore(finish)
	for _, node := range list[1:] {
		score := node.getScore(finish)
		if score < minScore {
			minScore = score
			min = node
		}
	}
	return min
}
