package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	t.Log("=========test IsPalindrome ===========")
	assert.Equal(t, true, IsPalindrome("wow"))
	assert.Equal(t, false, IsPalindrome("我爱你"))
	assert.Equal(t, false, IsPalindrome("柚子社区"))
	assert.Equal(t, true, IsPalindrome("tet"))

}
