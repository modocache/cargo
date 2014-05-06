package comparators_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestComparators(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Comparators Suite")
}
