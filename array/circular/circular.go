package circular

import "errors"

type Node struct {
	Value int
	next  *Node
}

type CircularBuffer struct {
	Size  int
	front *Node
	rear  *Node
	total int
}

func (cb *CircularBuffer) Insert(node *Node) {
	if cb.Size == 0 {
		panic(errors.New("Size not defined"))
	}
	if cb.total == cb.Size {
		panic(errors.New("Buffer is full"))
	}

	if cb.total == 0 {
		cb.front = node
		cb.rear = node
	} else {
		rear := cb.rear
		rear.next = node
		cb.rear = node
	}
	cb.total++
}

func (cb *CircularBuffer) Take() {
	if cb.total == 0 {
		panic(errors.New("Empty Buffer"))
	} else {
		cb.front = cb.front.next
	}
	cb.total--
}

func (cb *CircularBuffer) Length() int {
	return cb.total
}

func (cb *CircularBuffer) First() *Node {
	return cb.front
}

func (cb *CircularBuffer) Last() *Node {
	return cb.rear
}
