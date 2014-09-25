package main

import (
	"strconv"
	"strings"
)

//stock
type stock struct {
	data []int
	top  int
}

func NewStock(c int) *stock {
	data := make([]int, c)
	return &stock{data, 0}
}

func (s *stock) IsEmpty() bool {
	panicIfNil(s)
	return s.top == 0
}

func (s *stock) IsFull() bool {
	panicIfNil(s)
	return s.top == len(s.data)
}

func (s *stock) Len() int {
	return s.top
}

func (s *stock) Rest() int {
	panicIfNil(s)
	if s.IsFull() {
		return 0
	}
	return len(s.data) - s.top
}

func (s *stock) Push(iS ...int) {
	panicIfNil(s)
	if s.Rest() < len(iS) {
		new_data := make([]int, (len(s.data)+len(iS))*2)
		copy(new_data, s.data)
		s.data = new_data
	}

	for _, v := range iS {
		s.data[s.top] = v
		s.top++
	}
}

func (s *stock) Pop() (i int) {
	panicIfNil(s)
	s.top--
	i = s.data[s.top]
	return i
}

func (s *stock) search(card int) []int {
	if s.IsEmpty() {
		return nil
	}

	var cnt int
	for i := s.top - 1; i >= 0; i-- {
		if card == s.data[i] {
			break
		}
		cnt++
	}
	if cnt == s.top {
		return nil
	}

	sint := make([]int, cnt+1)
	for i := 0; i <= cnt; i++ {
		sint[i] = s.Pop()
	}
	return sint
}

func (s *stock) String() string {
	if s.top == 0 {
		return "empty"
	}
	slice := make([]string, s.top)
	for i := 0; i < s.top; i++ {
		slice[i] = strconv.Itoa(s.data[i])
	}

	return strings.Join(slice, " ")
}
