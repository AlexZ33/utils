package limit

import (
	"fmt"
	"sync"
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

func TestLimiter_Allow(t *testing.T) {
	// 限流器名称，发放令牌间隔，桶容量
	lmt := NewLimiter("xxxx", 1*time.Second, 20)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(10)
		go func() {
			if lmt.Allow() {
				fmt.Println("allow", i)
			} else {
				fmt.Println("not", i)
			}
		}()
	}
	wg.Wait()

}
