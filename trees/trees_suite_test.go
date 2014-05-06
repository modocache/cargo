package trees_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTrees(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Trees Suite")
}
