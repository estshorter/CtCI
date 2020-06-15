package p10

var root *RankNode

func track(number int) {
	if root == nil {
		root = new(number)
	} else {
		root.insert(number)
	}
}

func getRankOfNumber(number int) int {
	return root.getRank(number)
}

// RankNode defines...
type RankNode struct {
	Leftsize int
	Left     *RankNode
	Right    *RankNode
	data     int
}

func new(data int) *RankNode {
	return &RankNode{
		0, nil, nil, data,
	}
}

func (n *RankNode) insert(d int) {
	if d <= n.data {
		if n.Left != nil {
			n.Left.insert(d)
		} else {
			n.Left = new(d)
		}
		n.Leftsize++
	} else {
		if n.Right != nil {
			n.Right.insert(d)
		} else {
			n.Right = new(d)
		}
	}
}

func (n *RankNode) getRank(d int) int {
	if d == n.data {
		return n.Leftsize
	} else if d < n.data {
		if n.Left == nil {
			return -1
		}
		return n.Left.getRank(d)
	}
	rightRank := 0
	if n.Right == nil {
		rightRank = -1
	} else {
		rightRank = n.Right.getRank(d)
	}
	if rightRank == -1 {
		return -1
	}
	return n.Leftsize + 1 + rightRank
}
