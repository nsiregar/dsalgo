package twosum_test

import (
	"testing"

	. "github.com/nsiregar/dsalgo/array/twosum"
	"github.com/stretchr/testify/assert"
)

func TestTwoSumArray(t *testing.T) {
	arr := []int{2, 7, 11, 15}
	target := 9
	result := [2]int{2, 7}

	assert.Equal(t, result, TwoSum(arr, target))
}

func TestSortArray(t *testing.T) {
	arr := []int{1}
	target := 2

	assert.Panics(t, func() { TwoSum(arr, target) })
}
