package bubble_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestBubbleSorting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bubble Sorting Test Suite")
}
