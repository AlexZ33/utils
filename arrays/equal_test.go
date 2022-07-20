package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqual(t *testing.T) {
	t.Log("=========test Equal 深度相等判断 ===========")

	// 循环链表 Circular linked Lists  a->b->a and c->c
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	assert.Equal(t, true, Equal(a, a))  // true
	assert.Equal(t, true, Equal(b, b))  // true
	assert.Equal(t, true, Equal(c, c))  // true
	assert.Equal(t, false, Equal(a, b)) // false
	assert.Equal(t, false, Equal(b, c)) // false

	// output:
}
