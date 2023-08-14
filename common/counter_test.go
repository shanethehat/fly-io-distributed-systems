package common

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounterStartsAtZero(t *testing.T) {
	counter := NewCounter()
	assert.Equal(t, uint32(0), counter.count, "The value should be 0")
}

func TestCounterReturnsAnIncrementingInt(t *testing.T) {
	counter := NewCounter()
	for i := 1; i <= 3; i++ {
		assert.Equal(t, i, counter.IncrementAndRead(), fmt.Sprint("The incrementing counter should return ", i))
	}
}
