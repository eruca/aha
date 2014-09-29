package main

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := NewMap(5, 4)
	fmt.Println(m.data)
	if m == nil {
		t.Error("should not happened")
	}
	if m.x != 4 || m.y != 3 {
		t.Error("5,4", m.y, m.x)
	}

	if !m.Illegal(point{-1, 0}) || !m.Illegal(point{4, 3}) {
		t.Error("illegal but not")
	}

	if m.Arrive(point{2, 3}) {
		t.Error("illegal but not")
	}
	if !m.Arrive(point{3, 3}) {
		t.Error("illegal but not")
	}
}
