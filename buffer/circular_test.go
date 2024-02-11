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

import "testing"

func TestCicular(t *testing.T) {
	n := 10
	c := NewCircular[int](false, n)

	if !c.IsEmpty() {
		t.Fatalf("expected buffer to be empty, got: %v", c)
	}

	for i := 1; i <= n; i++ {
		if !c.Push(i) {
			t.Fatalf("unable to push value %v to circular buffer", i)
		}
	}

	if c.Push(n + 1) {
		t.Fatalf("expected buffer to be full")
	}

	var s int
	sum := func(n int) {
		s += n
	}
	c.Iter(sum)
	if exp := n * (n + 1) / 2; s != exp {
		t.Fatalf("expected sum of the buffer to be %v, got: %v", exp, s)
	}
}
