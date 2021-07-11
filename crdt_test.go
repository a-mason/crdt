package crdt

import (
	"fmt"
	"math"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	fmt.Println("Testing Counter")
	c1 := NewCounter()
	c1.Add("abc", 123)
	c1.Sub("abc", 120)
	c1.Add("bcd", 2)
	c1.Inc("bcd")
	c1.Dec("gdf")
	assert.Equal(t, int64(5), c1.Eval())
	assert.NotNil(t, c1.Add("abc", math.MaxInt64-122))
	assert.Equal(t, int64(5), c1.Eval())
	assert.NotNil(t, c1.Add("ddd", math.MaxInt64-2))
	assert.Equal(t, int64(5), c1.Eval())
	assert.NotNil(t, c1.Sub("abc", math.MaxInt64))
	assert.Equal(t, int64(5), c1.Eval())
	assert.Nil(t, c1.Add("ddd", math.MaxInt64-100))
	assert.Equal(t, int64(math.MaxInt64-95), c1.Eval())
}
