package hasdup_test

import (
	"testing"

	"github.com/nsiregar/dsalgo/array/hasdup"
	"github.com/stretchr/testify/assert"
)

func TestHasDupArrayRaisePanic(t *testing.T) {
	arr := []int{1}
	t.Log("raise panic")
	assert.Panics(t, func() { hasdup.Verify(arr) })
}

func TestHasDupArrayHasDuplicateMember(t *testing.T) {
	arr := []int{1, 1, 2}
	assert.True(t, hasdup.Verify(arr))
}

func TestHasDupArrayNoDuplicateMember(t *testing.T) {
	arr := []int{2, 1, 3}
	assert.False(t, hasdup.Verify(arr))
}

func BenchmarkHasDupArray(b *testing.B) {
	arr := []int{1, 1, 2, 3, 3, 4, 5, 7, 7, 8, 9, 9, 9, 0, 0, 1}
	for idx := 0; idx < b.N; idx++ {
		hasdup.Verify(arr)
	}
}
