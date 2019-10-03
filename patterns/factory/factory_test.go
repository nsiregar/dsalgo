package factory_test

import (
	"testing"

	"github.com/nsiregar/dsalgo/patterns/factory"
	"github.com/stretchr/testify/assert"
)

func TestFactoryPattern(t *testing.T) {
	baseStruct := factory.Base{
		Type:   "circle",
		Width:  20,
		Height: 0,
	}

	circle := factory.Factory(baseStruct)
	assert.IsType(t, factory.Circle{}, circle)

	baseStruct.Height = 20
	baseStruct.Type = "other"

	dummyShape := factory.Factory(baseStruct)
	assert.IsType(t, factory.DummyShape{}, dummyShape)
}
