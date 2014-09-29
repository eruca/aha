package main

import (
	"log"
	"strings"
)

type node struct {
	pt     point
	steps  int
	father int
}

var directions = [4][2]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

type Searcher struct {
	StartPt    point
	book       []point
	m          *Map
	queue      []node
	head, tail int
}

func NewSearcher(m *Map, pt point) *Searcher {
	book := make([]point, 1, 16)
	book[0] = pt
	queue := make([]node, 1, 16)
	queue[0].pt = pt

	return &Searcher{StartPt: pt, book: book, m: m, queue: queue, tail: 1}
}

func (s Searcher) IsWalked(pt point) bool {
	for _, v := range s.book {
		if v == pt {
			return true
		}
	}
	return false
}

func (s Searcher) Path() string {
	res := make([]string, 0, 16)
	for i := len(s.queue) - 1; i != 0; i = s.queue[i].father {
		res = append(res, s.queue[i].pt.String())
	}
	reverse := make([]string, len(res))
	var index int
	for i := len(res) - 1; i >= 0; i-- {
		reverse[index] = res[i]
		index++
	}
	return strings.Join(reverse, " ")
}

func (s *Searcher) Walk() string {
	var next point
Label:
	for s.head < s.tail {
		for _, dir := range directions {
			next = s.queue[s.head].pt.walk(dir)
			if s.m.Illegal(next) || s.IsWalked(next) {
				continue
			}
			s.book = append(s.book, next)

			var newNode node
			newNode.pt = next
			newNode.father = s.head
			newNode.steps = s.queue[s.head].steps + 1
			s.queue = append(s.queue, newNode)
			s.tail++

			if s.m.Arrive(next) {
				break Label
			}
		}
		s.head++
	}
	return s.Path()
}
