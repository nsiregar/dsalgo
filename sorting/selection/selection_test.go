package selection_test

import (
	"testing"

	"github.com/nsiregar/dsalgo/sorting/selection"
	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{5, 3, 1, 6, 7, 2, 4, 8}
	result := []int{1, 2, 3, 4, 5, 6, 7, 8}

	assert.Equal(t, selection.Sort(arr), result, "should sort array")
}
