package heap

import "testing"

func Test_MinHeap(t *testing.T) {
	heap := NewHeap(func(a, b int) bool { return a < b })

	heap.Push(1)
	heap.Push(4)
	heap.Push(3)
	heap.Push(2)
	heap.Push(1)

	if heap.Pop() != 1 {
		t.Error("Pop() failed")
	}

	if heap.Pop() != 1 {
		t.Error("Pop() failed")
	}

	if heap.Pop() != 2 {
		t.Error("Pop() failed")
	}

	if heap.Pop() != 3 {
		t.Error("Pop() failed")
	}

	if heap.Pop() != 4 {
		t.Error("Pop() failed")
	}
}

func Test_MaxHeap(t *testing.T) {
	heap := NewHeap(func(a, b int) bool { return a > b })

	heap.Push(1)
	heap.Push(2)
	heap.Push(3)
	heap.Push(4)

	if heap.Pop() != 4 {
		t.Error("Pop() failed")
	}

	if heap.Pop() != 3 {
		t.Error("Pop() failed")
	}

	if heap.Pop() != 2 {
		t.Error("Pop() failed")
	}
	if heap.Pop() != 1 {
		t.Error("Pop() failed")
	}
}

func Test_GenericHeap(t *testing.T) {
	type myType struct {
		value float32
	}

	heap := NewHeap(func(a, b myType) bool { return a.value < b.value })

	heap.Push(myType{value: 1})
	heap.Push(myType{value: 4})
	heap.Push(myType{value: 3})
	heap.Push(myType{value: 2})
	heap.Push(myType{value: 1})

	if heap.Pop().value != 1 {
		t.Error("Pop() failed")
	}

	if heap.Pop().value != 1 {
		t.Error("Pop() failed")
	}

	if heap.Pop().value != 2 {
		t.Error("Pop() failed")
	}

	if heap.Pop().value != 3 {
		t.Error("Pop() failed")
	}

	if heap.Pop().value != 4 {
		t.Error("Pop() failed")
	}
}
