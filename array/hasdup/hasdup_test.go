package hasdup_test

import (
	"testing"

	"github.com/nsiregar/dsalgo/array/hasdup"
	"github.com/stretchr/testify/assert"
)

func TestHasDupArray(t *testing.T) {
	t.Log("when array < 2")
	{
		arr := []int{1}
		t.Log("raise panic")
		{
			assert.Panics(t, func() { hasdup.Verify(arr) })
		}
	}

	t.Log("when array >= 2")
	{
		t.Log("\thas duplicate member")
		{
			arr := []int{1, 1, 2}
			t.Log("\t\treturns true")
			{
				assert.True(t, hasdup.Verify(arr))
			}
		}

		t.Log("\thas no duplicate member")
		{
			arr := []int{2, 1, 3}
			t.Log("\t\treturns false")
			{
				assert.False(t, hasdup.Verify(arr))
			}
		}
	}
}

func BenchmarkHasDupArray(b *testing.B) {
	arr := []int{1, 1, 2, 3, 3, 4, 5, 7, 7, 8, 9, 9, 9, 0, 0, 1}
	for idx := 0; idx < b.N; idx++ {
		hasdup.Verify(arr)
	}
}
