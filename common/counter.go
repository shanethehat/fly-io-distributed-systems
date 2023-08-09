package common

import (
	"sync/atomic"
)

type Counter struct {
	count uint64
}

func NewCounter() *Counter {
	return &Counter{count: 0}
}

func (counter *Counter) IncrementAndRead() int {
	atomic.AddUint64(&counter.count, 1)
	return int(counter.count)
}
