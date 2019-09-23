package circular_test

import (
	"testing"

	. "github.com/nsiregar/dsalgo/array/circular"
	"github.com/stretchr/testify/assert"
)

func TestCircularBuffer(t *testing.T) {
	node := &Node{
		Value: 10,
	}
	nodeTwo := &Node{
		Value: 12,
	}
	circular := CircularBuffer{
		Size: 10,
	}

	circular.Insert(node)
	assert.Equal(t, 1, circular.Length())
	assert.Equal(t, node, circular.First())

	circular.Insert(nodeTwo)
	assert.Equal(t, 2, circular.Length())
	assert.Equal(t, nodeTwo, circular.Last())
}

func TestFullCircularBuffer(t *testing.T) {
	node := &Node{
		Value: 10,
	}
	circular := CircularBuffer{
		Size: 5,
	}
	for idx := 0; idx < 5; idx++ {
		circular.Insert(node)
	}

	assert.Equal(t, 5, circular.Length())
	assert.Panics(t, func() { circular.Insert(node) })
}

func TestSizeNotDefined(t *testing.T) {
	node := &Node{
		Value: 10,
	}
	circular := new(CircularBuffer)

	assert.Panics(t, func() { circular.Insert(node) })
}

func TestEmptyCircularBuffer(t *testing.T) {
	circular := CircularBuffer{
		Size: 5,
	}

	assert.Panics(t, func() { circular.Take() })
}

func TestTakeCircularBuffer(t *testing.T) {
	node := &Node{
		Value: 10,
	}
	nodeTwo := &Node{
		Value: 12,
	}
	circular := CircularBuffer{
		Size: 10,
	}

	circular.Insert(node)
	circular.Insert(nodeTwo)
	circular.Take()
	assert.Equal(t, 1, circular.Length())
	assert.Equal(t, nodeTwo, circular.First())
	assert.Equal(t, nodeTwo, circular.Last())
}
