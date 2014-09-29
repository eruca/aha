package main

import (
	"fmt"
)

func main() {
	m := NewMap(5, 4)
	searcher := NewSearcher(m, point{0, 0})
	fmt.Println(searcher.Walk())
}
