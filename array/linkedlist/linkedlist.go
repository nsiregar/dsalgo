package linkedlist

import "errors"

type Item struct {
	Value int64
	Next  *Item
}

type LinkedList struct {
	Size int
	Head *Item
	Tail *Item
}

func (ll *LinkedList) Push(item *Item) {
	if ll.Size == 0 {
		ll.Head = item
		ll.Tail = item
	} else {
		tail := ll.Tail
		tail.Next = item
		ll.Tail = item
	}
	ll.Size++
}

func (ll *LinkedList) Pop() {
	if ll.Size == 0 {
		panic(errors.New("Empty List"))
	}
	var prevItem *Item
	currentItem := ll.Head
	for currentItem.Next != nil {
		prevItem = currentItem
		currentItem = currentItem.Next
	}
	prevItem.Next = currentItem.Next
	ll.Size--
}
