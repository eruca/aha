package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func (p point) walk(dir [2]int) (next point) {
	next.x = p.x + dir[0]
	next.y = p.y + dir[1]
	return next
}

func (p point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
