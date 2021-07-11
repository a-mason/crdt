package crdt

import (
	"errors"
	"math"
)

var ErrCounterOverflow = errors.New("counter overflow")
func sum(x int64, y int64) (sum int64, err error) {
		if x > math.MaxInt64 - y {
				return 0, ErrCounterOverflow
		}
    return x + y, nil
}

func maxInt64(x int64, y int64) (max int64) {
		if x > y {
				return x
		}
    return y
}
type Counter struct {
	countMap map[string][2]int64
}

func NewCounter() *Counter {
    var c Counter
    c.countMap = make(map[string][2]int64)
    return &c
}

func (c *Counter) Inc(id string) (err error) {
	return c.Add(id, 1);
}
func (c *Counter) Dec(id string) (err error) {
	return c.Sub(id, 1);
}

func (c *Counter) Add(id string, i int64) (err error) {
	if (i == 0) {
		return nil
	}
	cur := c.Eval()
	if (math.MaxInt64 - cur < i) {
		return ErrCounterOverflow
	}
	val := c.countMap[id]
	v, err := sum(val[0], i)
	if (err != nil) {
		return err
	}
	val[0] = v
	c.countMap[id] = val
	return nil
}

func (c *Counter) Sub(id string, i int64) (err error) {
	if (i == 0) {
		return nil
	}
	val := c.countMap[id]
	v, err := sum(val[1], i)
	if (err != nil) {
		return err
	}
	val[1] = v
	c.countMap[id] = val
	return nil
}

func (c Counter) Eval() (count int64) {
	for _, v := range c.countMap  {
		count += v[0]
		count -= v[1]
	}
	return
}

func Merge(c1 *Counter, c2 *Counter) (c *Counter) {
	c = new(Counter)
	// Copies c1 into c.
	for k, v := range c1.countMap {
		c.countMap[k] = [2]int64{v[0], v[1]}
	}
	// Merge by max
	for k, v := range c2.countMap {
		existing := c.countMap[k]
		existing[0] += maxInt64(existing[0], v[0])
		existing[1] += maxInt64(existing[1], v[1])
	}
	return
}
