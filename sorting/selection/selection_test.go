package selection_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nsiregar/dsalgo/sorting/selection"
)

var _ = Describe("Selection", func() {
	Describe("Sort", func() {
		It("sort by comparing", func() {
			arr := []int{5, 3, 1, 6, 7, 2, 4, 8}
			result := []int{1, 2, 3, 4, 5, 6, 7, 8}

			Expect(selection.Sort(arr)).To(Equal(result))
		})
	})
})
