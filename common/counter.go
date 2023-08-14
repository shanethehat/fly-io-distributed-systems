package common

import (
	"sync/atomic"
)

type Counter struct {
	count uint32
}

func NewCounter() *Counter {
	return &Counter{count: 0}
}

func (counter *Counter) IncrementAndRead() int {
	atomic.AddUint32(&counter.count, 1)
	return int(counter.count)
}
