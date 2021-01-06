package treap

import (
	"testing"
)

type Int int

func (i Int) Compare(other Comparable) int {
	oi := other.(Int)

	if int(i) < int(oi) {
		return -1
	} else if int(i) == int(oi) {
		return 0
	} else {
		return 1
	}
}

func randRange(min, max int) int {
	return min + rd.Int()%(max-min+1)
}

func TestRandKeyInsert(t *testing.T) {
	treap := New()
	for i := 1; i <= 10; i++ {
		treap.Insert(Int(randRange(100, 200)))
	}
	treap.checkBstInvariant()
	treap.checkHeapInvariant()
	t.Logf("\n%v\n", treap.String())
}

func TestSeqKeyInsert(t *testing.T) {
	treap := New()
	for i := 1; i <= 10; i++ {
		treap.Insert(Int(i))
	}
	treap.checkBstInvariant()
	treap.checkHeapInvariant()
	t.Logf("\n%v\n", treap.String())
}

func TestSearchMaxLE(t *testing.T) {
	treap := New()
	for i := 1; i <= 20; i = i + 2 {
		treap.Insert(Int(i))
	}
	treap.checkBstInvariant()
	treap.checkHeapInvariant()
	t.Logf("\n%v\n", treap.String())
	if e := treap.SearchMaxLE(Int(10)); e != Int(9) {
		t.Errorf("failed get max num <=  10")
	}
}

func TestExist(t *testing.T) {
	treap := New()
	for i := 1; i <= 20; i = i + 2 {
		treap.Insert(Int(i))
	}
	treap.checkBstInvariant()
	treap.checkHeapInvariant()
	t.Logf("\n%v\n", treap.String())
	if expect, ret := true, treap.Exist(Int(9)); expect != ret {
		t.Errorf("failed find key 9")
	}
	if expect, ret := false, treap.Exist(Int(18)); expect != ret {
		t.Errorf("shoud not find  key 18")
	}
}

func TestDelete(t *testing.T) {
	treap := New()
	for i := 1; i <= 10; i++ {
		treap.Insert(Int(i))
	}
	treap.Delete(Int(5))
	if expect, ret := false, treap.Exist(Int(5)); expect != ret {
		t.Errorf("failed delete key")
	}
	treap.checkBstInvariant()
	treap.checkHeapInvariant()

	t.Logf("\n%v\n", treap.String())
}
