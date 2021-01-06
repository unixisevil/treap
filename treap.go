package treap

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

const (
	LEFT = iota
	RIGHT
)

/*
  self < other, return negative value
  self > other, return positive value
  self == other,return zero value
*/
type Comparable interface {
	Compare(other Comparable) int
}

type node struct {
	links    [2]*node
	key      Comparable
	priority int
}

type Treap struct {
	root *node
}

var rd *rand.Rand

func init() {
	rd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func New() *Treap {
	return &Treap{}
}

func (n *node) printHelper(depth int, buf *bytes.Buffer) {
	if n == nil {
		return
	}

	n.links[LEFT].printHelper(depth+1, buf)

	for i := 0; i < depth; i++ {
		buf.WriteRune(' ')
	}
	buf.WriteString(fmt.Sprintf("%v [%d]\n", n.key, n.priority))

	n.links[RIGHT].printHelper(depth+1, buf)
}

func (t *Treap) String() string {
	var buf bytes.Buffer
	t.root.printHelper(0, &buf)
	return buf.String()
}

func (t *Treap) Exist(key Comparable) bool {
	w := t.root
	for w != nil {
		ret := key.Compare(w.key)
		switch {
		case ret == 0:
			return true
		case ret < 0:
			w = w.links[LEFT]
		case ret > 0:
			w = w.links[RIGHT]
		}
	}
	return false
}

/*
 return largest element <= key
 or nil if there is no such element.
*/

func (n *node) searchMaxLE(key Comparable) Comparable {
	if n == nil {
		return nil
	}
	if key.Compare(n.key) == 0 {
		return key
	} else if key.Compare(n.key) < 0 {
		return n.links[LEFT].searchMaxLE(key)
	} else {
		ret := n.links[RIGHT].searchMaxLE(key)
		if ret == nil {
			return n.key
		}
		return ret
	}
}

func (t *Treap) SearchMaxLE(key Comparable) Comparable {
	return t.root.searchMaxLE(key)
}

func (t *Treap) Insert(key Comparable) {
	insert(&t.root, key)
}

func (t *Treap) Delete(key Comparable) {
	delete(&t.root, key)
}

func (t *Treap) checkHeapInvariant() {
	t.root.checkHeapInvariant()
}

func (t *Treap) checkBstInvariant() {
	t.root.checkBstInvariant()
}

func (n *node) checkHeapInvariant() {
	if n == nil {
		return
	}
	n.links[LEFT].checkHeapInvariant()
	n.links[RIGHT].checkHeapInvariant()
	for dir := LEFT; dir <= RIGHT; dir++ {
		if c := n.links[dir]; c != nil && c.priority > n.priority {
			panic("heap invariant break")
		}
	}
}

func (n *node) checkBstInvariant() {
	if n == nil {
		return
	}
	n.links[LEFT].checkBstInvariant()
	n.links[RIGHT].checkBstInvariant()
	if c := n.links[LEFT]; c != nil && c.key.Compare(n.key) > 0 {
		panic("Bst invariant break")
	}
	if c := n.links[RIGHT]; c != nil && c.key.Compare(n.key) < 0 {
		panic("Bst invariant break")
	}
}

// sigle rotate at node  *pp
func rotateUp(pp **node, dir int) {
	if pp == nil {
		return
	}
	parent := *pp
	if parent == nil {
		return
	}
	child := parent.links[dir]
	if child == nil {
		return
	}
	grand := child.links[dir^1]
	*pp = child
	child.links[dir^1] = parent
	parent.links[dir] = grand
}

func insert(pp **node, key Comparable) {
	if *pp == nil {
		//no key
		*pp = &node{}
		(*pp).key = key
		(*pp).priority = rd.Int()
	} else if key.Compare((*pp).key) == 0 {
		return
	} else if key.Compare((*pp).key) < 0 {
		insert(&(*pp).links[LEFT], key)
	} else {
		insert(&(*pp).links[RIGHT], key)
	}
	//maintain heap property
	for dir := LEFT; dir <= RIGHT; dir++ {
		if (*pp).links[dir] != nil && (*pp).links[dir].priority > (*pp).priority {
			rotateUp(pp, dir)
		}
	}
}

func delete(pp **node, key Comparable) {
	if *pp == nil {
		return
	}
	if key.Compare((*pp).key) == 0 {
		for {
			//one child, or no child
			for dir := LEFT; dir <= RIGHT; dir++ {
				if (*pp).links[dir] == nil {
					*pp = (*pp).links[dir^1]
					return
				}
			}
			//two child
			var bigger int
			if (*pp).links[LEFT].priority > (*pp).links[RIGHT].priority {
				bigger = LEFT
			} else {
				bigger = RIGHT
			}
			rotateUp(pp, bigger)
			pp = &(*pp).links[bigger^1]
		}
	}

	var dir int
	if key.Compare((*pp).key) < 0 {
		dir = LEFT
	} else {
		dir = RIGHT
	}
	delete(&(*pp).links[dir], key)
}
