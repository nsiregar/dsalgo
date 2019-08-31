package bubble_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nsiregar/dsalgo/sorting/bubble"
)

var _ = Describe("Bubble", func() {
	Describe("Sort", func() {
		It("sorts array by sinking it", func() {
			arr := []int{5, 3, 1, 6, 7, 2, 4, 8}
			result := []int{1, 2, 3, 4, 5, 6, 7, 8}

			Expect(bubble.Sort(arr)).To(Equal(result))
		})
	})
})
