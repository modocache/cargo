package queues_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestQueues(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Queues Suite")
}
