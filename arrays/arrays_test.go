package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrays(t *testing.T) {
	t.Log("=========test arrays ===========")

	t.Log("test Compare function start")
	var arr = []interface{}{1, 2, 3, 4, nil, 5, "hello world", nil}

	t.Log(Compact(arr))
	result := Compact(arr)
	t.Log(result)
	assert.Equal(t, 6, len(result))
	assert.Equal(t, 1, result[0].(int))
	assert.Equal(t, 2, result[1].(int))
	assert.Equal(t, 3, result[2].(int))
	assert.Equal(t, 4, result[3].(int))
	assert.Equal(t, 5, result[4].(int))
	assert.Equal(t, "hello world", result[5].(string))
	//assert.Equal(t, nil, result[6])
	t.Log("test Compare function end")

	t.Log("test Flatten function start")
	var arr1 = []interface{}{1, 2, 3, 4}       // [1, 2, 3, 4]
	var arr2 = []interface{}{5, 6, 7, arr1}    // [5, 6, 7, [1, 2, 3, 4]]
	var arr3 = []interface{}{8, 9, arr1, arr2} // [8, 9, [1, 2, 3, 4], [5, 6, 7, [1, 2, 3, 4]]]

	r := Flatten(arr3)
	assert.Equal(t, 13, len(r))
	assert.Equal(t, 8, r[0].(int))
	assert.Equal(t, 9, r[1].(int))
	assert.Equal(t, 1, r[2].(int))
	assert.Equal(t, 2, r[3].(int))
	assert.Equal(t, 3, r[4].(int))
	assert.Equal(t, 4, r[5].(int))
	assert.Equal(t, 5, r[6].(int))
	assert.Equal(t, 6, r[7].(int))
	assert.Equal(t, 7, r[8].(int))
	assert.Equal(t, 1, r[9].(int))
	assert.Equal(t, 2, r[10].(int))
	assert.Equal(t, 3, r[11].(int))
	assert.Equal(t, 4, r[12].(int))

	t.Log("test Flatten function end")
}
