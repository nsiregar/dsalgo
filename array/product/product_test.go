package product_test

import (
	"testing"

	. "github.com/nsiregar/dsalgo/array/product"
	"github.com/stretchr/testify/assert"
)

func TestProductArray(t *testing.T) {
	arr := []int{10, 3, 5, 6, 2}
	result := []int{180, 600, 360, 300, 900}

	assert.Equal(t, result, Product(arr))
}
