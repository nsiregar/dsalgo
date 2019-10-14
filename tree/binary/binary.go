package binary

type Node struct {
	Value int64
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func NewNode(value int64) Node {
	node := Node{
		Value: value,
	}

	return node
}

func (n *Node) Insert(value int64) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &Node{
				Value: value,
			}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{
				Value: value,
			}
		} else {
			n.Right.Insert(value)
		}
	}
}

func (n *Node) Search(value int64) bool {
	if n == nil {
		return false
	}

	if n.Value > value {
		return n.Left.Search(value)
	}

	if n.Value < value {
		return n.Right.Search(value)
	}

	return true
}

func (t *Tree) Insert(value int64) *Tree {
	if t.Root == nil {
		t.Root = &Node{
			Value: value,
		}
	} else {
		t.Root.Insert(value)
	}

	return t
}

func (t *Tree) Search(value int64) bool {
	return t.Root.Search(value)
}
