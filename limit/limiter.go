/*
* Rerference
* 基于信号量的限流器：https://github.com/golang/net/blob/master/netutil/listen.go
* 滴滴开源了一个对 http 请求对限流器中间件：https://github.com/didip/tollbooth
* uber 开源了基于漏洞算法失效了一个限流器：https://github.com/uber-go/ratelimit Go 实现熔断器
* Go可用性(二) 令牌桶原理及使用:https://lailin.xyz/post/go-training-week6-2-token-bucket-1.html
 */

package limit

import (
	"context"
	"golang.org/x/time/rate"
	"strings"
	"sync"
	"time"
)

// 令牌桶算法限流 Traffic Shaping and Rate Limiting
//原理概述
//令牌：每次拿到令牌，才可访问
//桶 ，桶的最大容量是固定的，以固定的频率向桶内增加令牌，直至加满
//每个请求消耗一个令牌。
//限流器初始化的时候，令牌桶一般是满的。

type Limiter struct {
	name    string
	limiter *rate.Limiter
}

var m sync.Map

// NewLimiter interval 没隔多少时间添加一个令牌， 初始桶的容量
// // interval 表示放桶频率,是每秒放入的令牌数量
// burst 表示桶的大小 Limiter burst size
//构建一个限流器
func NewLimiter(name string, interval time.Duration, burst int) *Limiter {
	if lmt, ok := m.Load(strings.ToLower(name)); ok {
		return lmt.(*Limiter)
	}

	limiter := &Limiter{
		name:    name,
		limiter: rate.NewLimiter(rate.Every(interval), burst)}
	return limiter
}

// 使用时，每次都调用了 Allow() 方法
// Allow is shorthand for AllowN(time.Now(), 1).
// 就是 AllowN(now,1)  的别名
func (l *Limiter) Allow() bool {
	return l.limiter.Allow()
}

// AllowN reports whether n events may happen at time now.
// Use this method if you intend to drop / skip events that exceed the rate limit.
// Otherwise use Reserve or Wait.
// AllowN  表示截止到 now  这个时间点，是否存在 n  个 token，如果存在那么就返回 true 反之返回 false，如果我们限流比较严格，没有资源就直接丢弃可以使用这个方法。
func (l *Limiter) AllowN(now time.Time, n int) bool {
	return l.limiter.AllowN(now, n)
}

//Wait  是最常用的， Wait  是 WaitN(ctx, 1)  的别名
func (l *Limiter) Wait(ctx context.Context) error {
	return l.limiter.Wait(ctx)
}

//WaitN(ctx, n)  表示如果存在 n  个令牌就直接转发，不存在我们就等，等待存在为止，传入的 ctx  的 Deadline  就是等待的 Deadline
func (l *Limiter) WaitN(ctx context.Context, n int) error {
	return l.limiter.WaitN(ctx, n)
}
