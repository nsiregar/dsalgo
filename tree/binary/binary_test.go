package binary_test

import (
	"testing"

	"github.com/nsiregar/dsalgo/tree/binary"
	"github.com/stretchr/testify/assert"
)

var (
	root_value int64 = 24
)

func TestNodeCreation(t *testing.T) {
	node := binary.NewNode(12)
	result := binary.Node{
		Value: 12,
	}

	assert.Equal(t, node, result)
}

func TestInsertNodeLessOrEqual(t *testing.T) {
	var (
		first_value  int64 = 20
		second_value int64 = 20
		third_value  int64 = 18
	)
	node := binary.NewNode(root_value)

	node.Insert(first_value)
	assert.Equal(t, node.Left.Value, first_value)

	node.Insert(second_value)
	assert.Equal(t, node.Left.Left.Value, second_value)

	node.Insert(third_value)
	assert.Equal(t, node.Left.Left.Left.Value, third_value)
}

func TestInsertNodeGreater(t *testing.T) {
	var (
		first_value  int64 = 26
		second_value int64 = 28
		third_value  int64 = 30
	)

	node := binary.NewNode(root_value)

	node.Insert(first_value)
	assert.Equal(t, node.Right.Value, first_value)

	node.Insert(second_value)
	assert.Equal(t, node.Right.Right.Value, second_value)

	node.Insert(third_value)
	assert.Equal(t, node.Right.Right.Right.Value, third_value)
}

func TestBinaryTreeInsertion(t *testing.T) {
	var (
		first_value  int64 = 10
		second_value int64 = 15
	)
	binarytree := binary.Tree{}

	binarytree.Insert(root_value)
	assert.Equal(t, binarytree.Root.Value, root_value)

	binarytree.Insert(first_value)
	assert.Equal(t, binarytree.Root.Left.Value, first_value)

	binarytree.Insert(second_value)
	assert.Equal(t, binarytree.Root.Left.Right.Value, second_value)
}
