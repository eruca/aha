package main

import (
	"fmt"
	"math/rand"
	"time"
)

//player
type Player struct {
	q *queue
}

func NewPlayer(cards int) *Player {
	return &Player{NewQueue(cards)}
}

func (p *Player) Play() int {
	i, err := p.q.PopHead()
	if err != nil {
		panic(err)
	}
	return i
}

func (p *Player) Add(i ...int) {
	p.q.Push(i...)
}

func (p *Player) Empty() bool {
	return p.q.Len() == 0
}

func (p *Player) Card() string {
	return p.q.String()
}

func (p *Player) InitCard(card int) {
	p.q.Push(card)
}

//table
type Table struct {
	s *stock
}

func NewTable(c int) *Table {
	return &Table{NewStock(c)}
}

func (t *Table) Add(i int) (ret []int) {
	res := t.s.search(i)
	if res == nil {
		t.s.Push(i)
		return
	}
	ret = append(ret, i)
	ret = append(ret, res...)
	return ret
}

func (t *Table) Card() string {
	return t.s.String()
}

var A, B = NewPlayer(8), NewPlayer(8)

var T = NewTable(16)

//init
func Init() {
	rand.Seed(time.Now().UnixNano())
	var n int
	for i := 0; i < 8; i++ {
		n = rand.Intn(8) + 1
		A.InitCard(n)
	}
	for i := 0; i < 8; i++ {
		n = rand.Intn(8) + 1
		B.InitCard(n)
	}
	fmt.Println("A", A.Card())
	fmt.Println("B", B.Card())
}

func main() {
	Init()

	var card int
	var backs []int

	for !A.Empty() && !B.Empty() {
		card = A.Play()
		backs = T.Add(card)
		if backs != nil {
			A.Add(backs...)
		}

		card = B.Play()
		backs = T.Add(card)
		if backs != nil {
			B.Add(backs...)
		}
	}
	if A.Empty() {
		fmt.Println("A Win")
		fmt.Println("B cards:", B.Card())
	} else {
		fmt.Println("B Win")
		fmt.Println("A cards:", A.Card())
	}
	fmt.Println("Table cards:", T.Card())
}
