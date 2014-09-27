package arrayChain

import (
	"errors"
	"log"
	"unsafe"
)

type List struct {
	chain []*node
	head  int
	trail int
	exist map[unsafe.Pointer]struct{}
}

func NewList(l int) *List {
	chain := make([]*node, l)
	exist := make(map[unsafe.Pointer]struct{}, l)
	return &List{chain, -1, -1, exist}
}

func (this *List) Len() int {
	panicIfNil(this)
	return len(this.exist)
}

func (this *List) Cap() int {
	return len(this.chain)
}

//test just for prevent the same pt of node to add twice change the already one's index
//break
func (this *List) IsExist(n *node) (ok bool) {
	_, ok = this.exist[unsafe.Pointer(n)]
	return ok
}

func (this *List) Head() (n *node, pos int) {
	panicIfNil(this)
	return this.chain[this.head], this.head
}

func (this *List) Trail() (n *node, pos int) {
	panicIfNil(this)
	return this.chain[this.trail], this.trail
}

func (this *List) IsEmpty() bool {
	panicIfNil(this)
	return this.Len() == 0
}

func (this *List) IsFull() bool {
	panicIfNil(this)
	return len(this.chain) == this.Len()
}

func (this *List) pushIfEmpty(n *node) {
	this.head = 0
	this.trail = 0
	this.chain[0] = n
	this.chain[0].index = -1
	this.exist[unsafe.Pointer(n)] = struct{}{}
}

func (this *List) PushAfterPos(n *node, insertPos int) error {
	if insertPos < 0 || insertPos >= len(this.chain) {
		panic("pushPos:index out of range")
	}

	if this.chain[insertPos] == nil {
		panic("the node to been inserted after which is nil")
	}

	if this.IsExist(n) {
		return errors.New("the pt to the node is exist")
	}
	find, pos := this.searchNodeOfNil()
	if !find {
		this.pushIfFull(n, insertPos)
	}
	this.chain[pos] = n
	this.exist[unsafe.Pointer(n)] = struct{}{}
	this.pushAfterPos(pos, insertPos)
	return nil
}

func (this *List) pushAfterPos(nodePos, insertPos int) {
	this.chain[nodePos].index = this.chain[insertPos].index
	this.chain[insertPos].index = nodePos

	if insertPos == this.trail {
		this.trail = nodePos
	}
}

func (this *List) PopPos(popPos int) (n *node) {
	nodeBeforePos := this.getBeforePos(popPos)
	this.chain[nodeBeforePos].index = this.chain[popPos].index
	n = this.chain[popPos]
	this.chain[popPos] = nil
	delete(this.exist, unsafe.Pointer(n))

	return n
}

func (this *List) getBeforePos(pos int) (beforePos int) {
	if pos == this.head {
		return -1
	}

	var find bool
	for i := this.head; i != -1; {
		if i == pos {
			find = true
			break
		}
		beforePos = i
		i = this.chain[i].index
	}
	if !find {
		panic("the pos is not in the chain")
	}

	return beforePos
}

//insertPos if -1,insert head
func (this *List) pushIfFull(n *node, insertPos int) (head, next int) {
	if this.IsExist(n) {
		return -1, -1
	}

	newChain := make([]*node, len(this.chain)*2)

	var cnt int
	if insertPos == -1 {
		newChain[0] = n
		cnt++
		newChain[0].index = cnt
	}

	var tmp int
	for i := this.head; i != -1; {
		tmp = i

		newChain[cnt] = this.chain[i]
		i = this.chain[i].index
		//重塑index
		newChain[cnt].index = cnt + 1
		cnt++

		if tmp == insertPos {
			newChain[cnt] = n
			newChain[cnt].index = cnt + 1
			cnt++
		}
	}

	newChain[cnt-1].index = -1

	if cnt != len(this.chain)+1 {
		panic("the len is not correct")
	}

	this.chain = newChain
	this.exist[unsafe.Pointer(n)] = struct{}{}
	this.head = 0
	this.trail = cnt - 1

	return 0, cnt
}

func (this *List) searchNodeOfNil() (find bool, pos int) {
	if this.IsFull() {
		return false, -1
	}
	for pos = this.trail; pos < len(this.chain); pos++ {
		if this.chain[pos] == nil {
			find = true
			break
		}
	}
	if !find {
		for pos = 0; pos < this.trail; pos++ {
			if this.chain[pos] == nil {
				find = true
				break
			}
		}
	}
	if !find {
		return find, -1
	}
	return find, pos
}

func (this *List) PushTrail(n *node) error {
	panicIfNil(this)
	if this.IsExist(n) {
		return errors.New("the pt to the node is exist already!")
	}
	if this.IsEmpty() {
		this.pushIfEmpty(n)
		return nil
	}
	if this.IsFull() {
		this.pushIfFull(n, this.trail)
		return nil
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			this.pushIfFull(n, this.trail)
		}
	}()

	find, next := this.searchNodeOfNil()
	if !find {
		panic("should find one empty space")
	}
	this.chain[next] = n
	this.exist[unsafe.Pointer(n)] = struct{}{}

	this.pushAfterPos(next, this.trail)
	return nil
}

func (this *List) PushHead(n *node) {
	panicIfNil(this)
	if this.IsEmpty() {
		this.PushTrail(n)
		return
	}

	if this.IsFull() {
		this.pushIfFull(n, -1)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			this.pushIfFull(n, -1)
		}
	}()

	find, next := this.searchNodeOfNil()
	if !find {
		panic("should find one empty space")
	}

	this.chain[next] = n
	this.exist[unsafe.Pointer(n)] = struct{}{}

	this.chain[next].index = this.head
	this.head = next
}

func (this *List) PopHead() (n *node, pos int, err error) {
	panicIfNil(this)
	if this.IsEmpty() {
		return nil, -1, errors.New("the list is empty")
	}
	n = this.chain[this.head]
	pos = this.head

	this.chain[this.head] = nil
	delete(this.exist, unsafe.Pointer(n))

	this.head = n.index

	return n, pos, nil
}

func (this *List) PopTrail() (n *node, pos int, err error) {
	panicIfNil(this)
	if this.IsEmpty() {
		return nil, -1, errors.New("the list is empty")
	}
	n = this.chain[this.trail]
	pos = this.trail

	this.chain[this.trail] = nil
	delete(this.exist, unsafe.Pointer(n))

	this.trail = this.getBeforePos(this.trail)

	return n, pos, nil
}
