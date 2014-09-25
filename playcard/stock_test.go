package main

import (
	// "fmt"
	"testing"
)

var s = NewStock(4)

func TestNewStock(t *testing.T) {
	if s == nil {
		t.Error("stock not nil, but nil")
	}
}

func TestStockPush(t *testing.T) {
	if !s.IsEmpty() {
		t.Error("empty,but not")
	}

	s.Push(1, 2, 3)
	if s.Len() != 3 {
		t.Error("len 3", s.Len())
	}

	if s.Rest() != 1 {
		t.Error("rest 5", s.Rest())
	}
	s.Push(1)

	if !s.IsFull() {
		t.Error("full,but not")
	}

	if s.String() != "1 2 3 1" {
		t.Error("should 1 2 3 1,but not")
	}
	s.search(2)
	//fmt.Println("1 3 2", ints)

	s.Pop()
	if !s.IsEmpty() {
		t.Error("empty,but not")
	}
}
