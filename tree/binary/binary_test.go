package binary_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nsiregar/dsalgo/tree/binary"
)

var _ = Describe("Binary Tree", func() {
	Describe("Node", func() {
		Describe("#NewNode", func() {
			It("initialize node object", func() {
				node := binary.NewNode(12)
				result := binary.Node{
					Value: 12,
				}

				Expect(node).To(Equal(result))
			})
		})

		Describe("#Insert", func() {
			Context("less or equal than current node value", func() {
				Context("left node nil", func() {
					var new_value int64

					It("insert to the left", func() {
						node := binary.NewNode(24)
						new_value = 20

						node.Insert(new_value)

						Expect(node.Left.Value).To(Equal(new_value))
					})
				})

				Context("left node not nil", func() {
					var second_node int64
					var third_node int64

					It("call #Insert on nodes", func() {
						node := binary.NewNode(24)
						second_node = 20
						third_node = 18

						node.Insert(second_node)
						node.Insert(third_node)

						Expect(node.Left.Value).To(Equal(second_node))
						Expect(node.Left.Left.Value).To(Equal(third_node))
					})
				})
			})

			Context("greater than current node value", func() {
				Context("right node nil", func() {
					var new_value int64

					It("insert to the right", func() {
						node := binary.NewNode(24)
						new_value = 25

						node.Insert(new_value)

						Expect(node.Right.Value).To(Equal(new_value))
					})
				})

				Context("right node not nil", func() {
					var second_node int64
					var third_node int64

					It("call #Insert on nodes", func() {
						node := binary.NewNode(24)
						second_node = 25
						third_node = 30

						node.Insert(second_node)
						node.Insert(third_node)

						Expect(node.Right.Value).To(Equal(second_node))
						Expect(node.Right.Right.Value).To(Equal(third_node))
					})
				})
			})
		})
	})

	Describe("BinaryTree", func() {
		Context("#Insert", func() {
			Context("root is nil", func() {
				var root_value int64

				It("create root node", func() {
					binarytree := binary.Tree{}
					root_value = 64

					binarytree.Insert(root_value)

					Expect(binarytree.Root.Value).To(Equal(root_value))
				})
			})

			Context("root is not nil", func() {
				It("insert to nodes", func() {
					binarytree := binary.Tree{}
					binarytree.Insert(10)
					binarytree.Insert(15)

					Expect(binarytree.Root.Value).To(Equal(int64(10)))
					Expect(binarytree.Root.Right.Value).To(Equal(int64(15)))
				})
			})
		})
	})
})
