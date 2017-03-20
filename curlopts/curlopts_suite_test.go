package curlopts_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCurlopts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Curlopts Suite")
}
