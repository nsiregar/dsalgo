package bubble_test

import (
	"testing"

	"github.com/nsiregar/dsalgo/sorting/bubble"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	t.Log("should sort array")
	{
		arr := []int{5, 3, 1, 6, 7, 2, 4, 8}
		result := []int{1, 2, 3, 4, 5, 6, 7, 8}

		assert.Equal(t, result, bubble.Sort(arr))
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	arr := []int{5, 3, 1, 6, 7, 2, 4, 8}

	for idx := 0; idx < b.N; idx++ {
		bubble.Sort(arr)
	}
}
