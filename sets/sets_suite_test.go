package sets_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSets(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sets Suite")
}
