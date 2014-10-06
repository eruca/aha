package pipe

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := NewMap()
	m.Init()
	m.Start()
	if m.find {
		fmt.Println("find")
	} else {
		fmt.Println("failed")
	}
}
