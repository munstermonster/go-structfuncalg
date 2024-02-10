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

type Node[T any] struct {
	val  T
	next *Node[T]
}
<<<<<<< HEAD
=======

func (n *Node[T]) SetVal(v T) {
	n.val = v
}

func (n *Node[T]) GetVal() T {
	return n.val
}

func (n *Node[T]) SetNext(next *Node[T]) {
	n.next = next
}

func (n *Node[T]) Next() *Node[T] {
	if n == nil {
		return nil
	}
	return n.next
}
>>>>>>> 190e7a2 (linkedlist: Add package)
