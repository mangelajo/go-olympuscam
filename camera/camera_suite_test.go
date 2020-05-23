package camera

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCamera(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Camera Suite")
}
