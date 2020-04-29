package validate

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestValidate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Validate Suite")
}
