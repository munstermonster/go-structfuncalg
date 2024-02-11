// Copyright (c) 2024 Andrew Gloster

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of c software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and c permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package buffer

import "fmt"

// Circular implements a cicular buffer of type T where Push and Pop are FIFO
type Circular[T any] struct {
	overwrite bool
	count     int
	head      *int
	tail      *int
	buf       []*T
}

func (c Circular[T]) String() string {
	var s string
	for _, v := range c.buf {
		if v == nil {
			s = s + fmt.Sprintf("nil, ")
		} else {
			s = s + fmt.Sprintf("%v, ", *v)
		}
	}
	return fmt.Sprintf("%v, %v, %v", *c.head, *c.tail, s)
}

func NewCircular[T any](overwrite bool, n int) Circular[T] {
	return Circular[T]{overwrite: overwrite, buf: make([]*T, n)}
}

func (c *Circular[T]) Push(value T) bool {
	if c.IsFull() {
		return false
	}
	if c.count != len(c.buf) {
		c.count++
	}
	if c.IsEmpty() {
		var a, b int
		c.head = &a
		c.tail = &b
		c.buf[0] = &value
		return true
	}
	idx := (*c.tail + 1) % len(c.buf)
	c.buf[idx] = &value
	c.tail = &idx
	return true
}

func (c *Circular[T]) Pop() (bool, T) {
	if c.IsEmpty() {
		return false, *new(T)
	}
	v := *c.buf[*c.head]
	idx := (*c.head + 1) % len(c.buf)
	if c.buf[idx] != nil {
		c.head = &idx
	} else {
		c.head = nil
		c.tail = nil
	}
	c.count--
	return true, v
}

func (c *Circular[T]) Iter(f func(T)) {
	if c.IsEmpty() {
		return
	}
	for _, v := range c.buf {
		if v == nil {
			continue
		}
		f(*v)
	}
}

func (c *Circular[T]) Front() T {
	return *c.buf[*c.head]
}

func (c *Circular[T]) Rear() T {
	return *c.buf[*c.tail]
}

func (c *Circular[T]) Size() int {
	return c.count
}

func (c *Circular[T]) IsEmpty() bool {
	return c.head == nil && c.tail == nil
}

func (c *Circular[T]) IsFull() bool {
	return (*c.tail+1)%len(c.buf) == *c.head && !c.overwrite
}
