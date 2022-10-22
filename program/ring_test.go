package program

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRingArray_Add(t *testing.T) {
	r := RingArray{Values: make([]int, 0, 5)}

	r.Add(2)
	assert.Equal(t, 2, r.Average())
	r.Add(4)
	assert.Equal(t, 3, r.Average())
	r.Add(4)
	assert.Equal(t, 3, r.Average())
	r.Add(1)
	assert.Equal(t, 2, r.Average())
	r.Add(10)
	assert.Equal(t, 4, r.Average())

	r.Add(10)
	assert.Equal(t, 5, r.Average())

	r.Add(10)
	assert.Equal(t, 7, r.Average())
	r.Add(10)
	assert.Equal(t, 8, r.Average())
	r.Add(10)
	assert.Equal(t, 10, r.Average())

}
