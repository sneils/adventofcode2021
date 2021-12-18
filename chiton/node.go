package chiton

type Node struct {
	x, y, risk int
	// g = this.getRisk()
	// h = this.getDistance(finish)
	// f = g + h
	parent *Node
}

func (node Node) getRisk() int {
	risk := 0
	cur := node
	for cur.parent != nil {
		risk += cur.risk
		cur = *cur.parent
	}
	return risk
}

func (node Node) getDistance(other Node) int {
	xd, yd := node.x-other.x, node.y-other.y
	if xd < 0 {
		xd *= -1
	}
	if yd < 0 {
		yd *= -1
	}
	return xd + yd
}

func (node Node) getScore(finish Node) int {
	return node.getRisk() + node.getDistance(finish)
}
