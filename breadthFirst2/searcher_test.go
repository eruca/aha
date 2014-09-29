package main

import (
	"fmt"
	"testing"
)

func TestSearcher(t *testing.T) {
	m := NewMap(5, 4)
	s := NewSearcher(m, point{0, 0})
	fmt.Println(s.Walk())
}
