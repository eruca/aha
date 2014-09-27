package arrayChain

import (
	"testing"
)

func expect(t *testing.T, lhs, rhs int) {
	if lhs != rhs {
		t.Errorf("should be %d,but %d", lhs, rhs)
	}
}

func TestNewNode(t *testing.T) {
	n := NewNode(1)
	expect(t, n.index, -1)
}
