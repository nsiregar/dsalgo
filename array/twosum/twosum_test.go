package twosum_test

import (
	"testing"

	. "github.com/nsiregar/dsalgo/array/twosum"
	"github.com/stretchr/testify/assert"
)

func TestTwoSumArray(t *testing.T) {
	t.Log("when length >= 2")
	{
		t.Logf("\tfind two value which sum is target")
		{
			arr := []int{2, 7, 11, 15}
			target := 9
			result := [2]int{2, 7}

			assert.Equal(t, result, TwoSum(arr, target))
		}

		t.Log("\twhen value not found")
		{
			t.Log("\t\treturn array of 0")
			{
				arr := []int{4, 7, 11, 15}
				target := 9
				result := [2]int{0, 0}

				assert.Equal(t, result, TwoSum(arr, target))
			}
		}
	}

	t.Log("when length < 2")
	{
		t.Logf("\traise panic when array length < 2")
		{
			arr := []int{1}
			target := 2

			assert.Panics(t, func() { TwoSum(arr, target) })
		}
	}
}
