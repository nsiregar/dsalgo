package linkedlist_test

import (
	"testing"

	. "github.com/nsiregar/dsalgo/array/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestLinkedListBehavior(t *testing.T) {
	var emptyItem *Item
	firstItem := &Item{
		Value: 10,
	}
	secondItem := &Item{
		Value: 12,
	}
	linkedlist := new(LinkedList)

	linkedlist.Push(firstItem)
	assert.Equal(t, linkedlist.Size, 1)
	assert.Equal(t, linkedlist.Head.Value, firstItem.Value)

	linkedlist.Push(secondItem)
	assert.Equal(t, linkedlist.Size, 2)
	assert.Equal(t, linkedlist.Head.Next.Value, secondItem.Value)

	linkedlist.Pop()
	assert.Equal(t, linkedlist.Size, 1)
	assert.Equal(t, linkedlist.Head.Value, firstItem.Value)
	assert.Equal(t, linkedlist.Head.Next, emptyItem)
}
