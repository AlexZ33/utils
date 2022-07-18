package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Log("test arrays")
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
	assert.Equal(t, nil, result[6])
}
