package linkedlist

import "fmt"

type Item struct {
	Value int64
	Next  *Item
}

type LinkedList struct {
	Size int
	Head *Item
}

func (ll *LinkedList) Push(item *Item) {
	if ll.Size == 0 {
		ll.Head = item
	} else {
		currentItem := ll.Head
		for currentItem.Next != nil {
			currentItem = currentItem.Next
		}
		currentItem.Next = item
	}
	ll.Size++
}

func (ll *LinkedList) Pop() {
	if ll.Size == 0 {
		fmt.Printf("Empty List")
	} else {
		currentItem := ll.Head
		prevItem := ll.Head
		for currentItem.Next != nil {
			prevItem = currentItem
			currentItem = currentItem.Next
		}
		prevItem.Next = nil
	}
	ll.Size--
}
