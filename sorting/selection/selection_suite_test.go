package selection_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSelectionSorting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Selection Sorting Test Suite")
}
