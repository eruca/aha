package pipe

import (
	"log"
)

type Pipe struct {
	Type   state
	rawDir int
	book   bool
}

func NewPipe(dir int) *Pipe {
	var t state
	switch dir {
	case 0, 1, 2, 3:
		t = pipeL
	case 4, 5:
		t = pipeI
	default:
		log.Panicln(dir, "should not happened")
	}
	return &Pipe{Type: t, rawDir: dir}
}
