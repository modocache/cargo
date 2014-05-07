package graphs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGraphs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Graphs Suite")
}
