package main

import (
	// "fmt"
	"testing"
)

var q = NewQueue(8)

func TestLen(t *testing.T) {
	if q.Len() != 0 {
		t.Error("0,but", q.Len())
	}
	q.Push(1, 2, 3, 4)
	if q.Len() != 4 {
		t.Error("4,but", q.Len())
	}
	if q.Rest() != 4 {
		t.Error("4,but", q.Rest())
	}

	q.Push(5, 6, 7, 8, 9)
	if q.Len() != 9 {
		t.Error("9 but", q.Len())
	}
}

func TestPop(t *testing.T) {
	i, err := q.PopHead()
	if err != nil {
		t.Error(err)
	}
	if i != 1 {
		t.Error("1,but", i)
	}

	q.PopHead()
	q.PopHead()
	if q.Len() != 6 {
		t.Error("6 but", q.Len())
	}
	if q.String() != "4 5 6 7 8 9" {
		t.Error("4 5 6 7 8 9--", q.String())
	}
}

func TestString(t *testing.T) {
	que := NewQueue(4)
	que.Push(1, 2, 3, 4)
	que.PopHead()
	que.PopHead()
	que.PopHead()
	que.Push(1, 2)
	if que.String() != "4 1 2" {
		t.Error("4 1 2,but", que.String())
	}
}
