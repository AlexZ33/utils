package lrucache

import (
	"container/list"
	"testing"
)

type Elem struct {
	key   int
	value string
}

func Test_New(t *testing.T) {
	lc := New(5)
	if lc.Len() != 0 {
		t.Error("case 1 failed")
	}
}

func Test_Put(t *testing.T) {
	lc := New(0)
	lc.Put(1, "1")
	if lc.Len() != 0 {
		t.Error("capacity 0 : case 1.1 failed")
	}

	lc = New(5)
	lc.Put(1, "1")
	lc.Put(1, "2")
	lc.Put(1, "3")

	if lc.Len() != 2 {
		t.Error("capacity 5: case 2.1 failed")
	}

	l := list.New()
	l.PushBack(&Elem{1, "3"})
	l.PushBack(&Elem{2, "2"})
	e := l.Front()
	for c := lc.Front(); c != nil; c = c.Next() {
		v := e.Value.(*Elem)
		if c.Key.(int) != v.key {
			t.Error("case 2.2 failed", c.Key.(int), v.key)
		}
		if c.Value.(string) != v.value {
			t.Error("case 2.3 failed", c.Value.(string), v.value)
		}
		e = e.Next()
	}

	lc.Put(3, "4")
	lc.Put(4, "5")
	lc.Put(5, "6")
	lc.Put(2, "7")
	if lc.Len() != 5 {
		t.Error("case 3.1 failed")
	}

}
