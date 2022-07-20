// twitter雪花算法golang实现, 生成唯一趋势自增id

package algorithm

import "sync"

type SnowFlake struct {
	_lock     sync.Mutex // 锁
	timestamp int64      // 时间戳
	sequence  int64      // 序列号占12位，十进制位范围是[0, 4095]
}
