package main

type Map struct {
	data [][]int
	x    int
	y    int
}

func NewMap(n, m int) *Map {
	var M Map
	M.data = make([][]int, n)
	for k, _ := range M.data {
		M.data[k] = make([]int, m)
	}

	M.x = n - 1
	M.y = m - 1
	M.Init()
	return &M
}

func (m *Map) Init() {
	m.data[0][2] = -1
	m.data[2][2] = -1
	m.data[3][1] = -1
	m.data[4][3] = -1
	m.data[3][3] = 1
}

func (m Map) find(pt point) int {
	return m.data[pt.x][pt.y]
}

func (m Map) Illegal(pt point) bool {
	return pt.x < 0 || pt.x > m.x || pt.y < 0 ||
		pt.y > m.y || m.find(pt) == -1
}

func (m Map) Arrive(pt point) bool {
	return m.find(pt) == 1
}
