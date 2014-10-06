package arrayChain

import (
	// "fmt"
	"testing"
)

func TestNewList(t *testing.T) {
	list := NewList(10)

	expect(t, list.Cap(), 10)
	expect(t, list.Len(), 0)
	if !list.IsEmpty() {
		t.Error("empty,but not")
	}

	if list.IsFull() {
		t.Error("not full,but full")
	}
}

func TestHeadTail(t *testing.T) {
	list := NewList(10)
	list.PushHead(NewNode(0))
	h, hpos := list.Head()
	tr, tpos := list.Tail()
	if h != tr || hpos != tpos {
		t.Error("head equal,but not")
	}

	if h.data != 0 {
		t.Error("0,but", h.data)
	}

	if list.Len() != 1 {
		expect(t, 1, list.Len())
	}
}

func TestpushIfEmpty(t *testing.T) {
	list := NewList(10)
	if list.IsEmpty() {
		list.pushIfEmpty(NewNode(0))
	}
	n, index := list.Head()
	if index != -1 || n.data != 0 {
		t.Error("TestpushIfEmpty true ,but false")
	}
}

func TestPos(t *testing.T) {
	list := NewList(10)
	list.PushTail(NewNode(0))
	list.PushTail(NewNode(1))
	list.PushTail(NewNode(2))
	list.PushTail(NewNode(3))
	if list.Len() != 4 {
		expect(t, 4, list.Len())
	}
	for i := 0; i < 3; i++ {
		if list.chain[i].index != i+1 {
			t.Error("list.chain[i].index != i+1")
		}
	}
	if list.chain[3].index != -1 {
		expect(t, -1, list.chain[3].index)
	}

	if list.getBeforePos(list.head) != -1 || list.getBeforePos(list.tail) != 2 {
		t.Error("true,but false")
	}
}

func TestPos2(t *testing.T) {
	list := NewList(5)
	list.PushTail(NewNode(0))
	list.PushTail(NewNode(1))
	list.PushTail(NewNode(2))
	list.PushTail(NewNode(3))
	list.PushTail(NewNode(4))
	list.PushTail(NewNode(5))
	if list.Cap() != 10 {
		t.Error("10,but", list.Cap())
	}

	n := list.PopPos(3)
	if n.index != 4 {
		t.Error("2,but", n.index)
	}

	list.PopHead()
	list.PopTail()
	expect(t, 3, list.Len())
	list.PushHead(n)
	expect(t, 4, list.Len())
	head, _ := list.Head()
	expect(t, 3, head.data)
	expect(t, 1, head.index)
}

func TestPos3(t *testing.T) {
	list := NewList(5)
	list.PushTail(NewNode(0))
	list.PushTail(NewNode(1))
	list.PushTail(NewNode(2))
	list.PushHead(NewNode(-1))
	list.PushAfterPos(NewNode(3), 3)
	list.PopHead()
	list.PopTail()
	list.PopTail()
	list.PopTail()
	list.PopHead()
	if list.Len() != 0 {
		t.Error("0,but", list.Len())
	}
}
