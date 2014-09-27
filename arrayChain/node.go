package arrayChain

type node struct {
	data  int
	index int
}

func NewNode(data int) *node {
	return &node{data, -1}
}
