package limit

import (
	"fmt"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	limiter := NewLimiter("test", time.Millisecond*500, 2)
	//time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		var ok bool
		if limiter.Allow() {
			ok = true
		}
		time.Sleep(time.Millisecond * 20)
		fmt.Println(ok)
	}
}
