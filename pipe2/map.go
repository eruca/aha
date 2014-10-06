package pipe

// import (
// 	"fmt"
// 	// "math/rand"
// 	// "time"
// )

//to record the ways,there are maybe several ways
type Records struct {
	x      int
	y      int
	cnt    int
	father int
}

var res = []int{4, 0, 1, 5, 0, 2, 4, 2, 0, 4, 4, 2, 4, 0, 4, 2, -1, 0, 4, 3}

//finds index the end point,and the ways be the arraychain
type Map struct {
	M [][]*Pipe

	inDir, outDir int
	maxX, maxY    int

	Ways  []Records
	finds []int //index of the find in Ways
}

func NewMap() *Map {
	M := make([][]*Pipe, 4)
	var index int
	for i := 0; i < 4; i++ {
		M[i] = make([]*Pipe, 5)
		for j := 0; j < 5; j++ {
			if i == 3 && j == 1 {
				index++
				continue
			}
			M[i][j] = NewPipe(res[index])
			// fmt.Printf("%d ", res[index])
			index++
		}
		// fmt.Println("")
	}
	return &Map{M: M, maxX: 3, maxY: 4}
}

// func NewMap(n, m int) *Map {
// 	time.Now().UnixNano()
// 	M := make([][]*Pipe, m)
// 	for k, _ := range M {
// 		M[k] = make([]*Pipe, n)
// 		for i := 0; i < n; i++ {
// 			if k == 1 && i == 3 {
// 				continue
// 			}
// 			M[k][i] = NewPipe(rand.Intn(6))
// 			fmt.Print(M[k][i].Type, " ")
// 		}
// 		fmt.Println("")
// 	}

// 	ways := make([]Records, 0, 16)

// 	return &Map{M: M, maxX: n - 1, maxY: m - 1, Ways: ways}
// }

func (m *Map) Init() {
	m.inDir = 2
	m.outDir = 2
}

func (m *Map) Start() {
	m.M[0][0].book = true
	switch m.M[0][0].Type {
	case pipeI:
		cnt := m.M[0][0].rawDir % 2
		r := Records{x: 0, y: 0, cnt: cnt}
		m.Ways = append(m.Ways, r)
		m.Next(1, 0, m.inDir)
	case pipeL:
		cnt := (5 - m.M[0][0].rawDir) &^ 4
		m.Ways = append(m.Ways, Records{0, 0, cnt, 0})
		m.Next(0, 1, m.inDir-1)
	}
	m.M[0][0].book = false
}

func (m *Map) Next(x, y, front int) {
	if x == m.maxX+1 && y == m.maxY && front == m.outDir {
		m.find = true
		println("find way to finish it")
		return
	}

	if x < 0 || x > m.maxX || y < 0 || y > m.maxY || m.M[x][y] == nil || m.M[x][y].book {
		return
	}
	// fmt.Printf("x:%d,y:%d,front:%d\n", x, y, front)

	m.M[x][y].book = true

	switch m.M[x][y].Type {
	case pipeI:
		switch front {
		case 0:
			m.Next(x-1, y, front)
		case 1:
			m.Next(x, y-1, front)
		case 2:
			m.Next(x+1, y, front)
		case 3:
			m.Next(x, y+1, front)
		}
	case pipeL:
		switch front {
		case 0, 2:
			m.Next(x, y-1, 1)
			m.Next(x, y+1, 3)
		case 1, 3:
			m.Next(x-1, y, 0)
			m.Next(x+1, y, 2)
		}
	}
	m.M[x][y].book = false
}
