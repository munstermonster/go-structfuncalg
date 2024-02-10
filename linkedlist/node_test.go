// MIT License

// Copyright (c) 2024 Andrew Gloster

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package linkedlist

import (
	"testing"
)

func TestNode(t *testing.T) {
	num := 100
	nodes := make([]Node[int], num)
	var i int
	for i < num {
		nodes[i].SetVal(i)
		if i < num-1 {
			nodes[i].SetNext(&nodes[i+1])
		}
		i++
	}

	var exp int
	cur := &nodes[0]
	for cur != nil {
		if v := cur.GetVal(); v != exp {
			t.Fatalf("expected value of node to be %v, got: %v", exp, v)
		}
		exp++
		cur = cur.Next()
	}
}
