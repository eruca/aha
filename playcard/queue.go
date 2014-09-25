package main

import (
	"strconv"
	"strings"
)

//queue
type queue struct {
	data  []int
	head  int
	trail int
	cross bool
}

func NewQueue(c int) *queue {
	return &queue{data: make([]int, c)}
}

func (q *queue) Len() int {
	if !q.cross && q.trail >= q.head {
		return q.trail - q.head
	}
	if q.cross && q.head >= q.trail {
		return len(q.data[:q.trail]) + len(q.data[q.head:])
	}
	return 0
}

func (q *queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *queue) Rest() int {
	return len(q.data) - q.Len()
}

func (q *queue) Push(i ...int) {
	var data []int
	if q.Rest() < len(i) {
		data = make([]int, (len(q.data)+len(i))*2)
		if !q.cross {
			copy(data, q.data[q.head:q.trail])
		} else {
			copy(data, q.data[q.head:])
			copy(data[len(q.data)-q.head:], q.data[:q.trail])
		}
		q.trail = q.Len()
		q.head = 0
		q.data = data
		q.cross = false
	}
	for _, v := range i {
		q.data[q.trail] = v
		q.trail++
		if q.trail == len(q.data) {
			q.trail = 0
			q.cross = true
		}
	}
}

func (q *queue) PopHead() (i int, err error) {
	panicIfNil(q)

	i = q.data[q.head]
	q.head++
	if q.head == len(q.data) && q.cross {
		q.cross = false
		q.head = 0
	}
	return i, nil
}

func (q *queue) String() string {
	slice := make([]string, q.Len())
	if !q.cross {
		for k, v := range q.data[q.head:q.trail] {
			slice[k] = strconv.Itoa(v)
		}
	} else {
		var k, v int
		for k, v = range q.data[q.head:] {
			slice[k] = strconv.Itoa(v)
		}

		for i, j := range q.data[:q.trail] {
			slice[k+i+1] = strconv.Itoa(j)
		}
	}
	return strings.Join(slice, " ")
}
