package curlopts_test

import (
	. "github.com/concourse/archive-resource/curlopts"
	"github.com/concourse/archive-resource/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CurlOpts", func() {

	var (
		source   models.Source
		curlOpts CurlOpts
		err      error
	)

	JustBeforeEach(func() {
		curlOpts, err = NewCurlOpts(source)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("String()", func() {

		var resultingArgs string

		JustBeforeEach(func() {
			resultingArgs = curlOpts.String()
		})

		Context("when the source contains a CA Cert", func() {
			BeforeEach(func() {
				source = models.Source{CaCert: "blah"}
			})

			It("returns a '--cacert' flag with the cert path", func() {
				Expect(resultingArgs).To(MatchRegexp("--cacert [\\/a-zA-Z0-9]*archive-resource-cert[0-9]*"))
			})
		})

		Context("when the source specifies SkipSslValidation", func() {
			BeforeEach(func() {
				source = models.Source{SkipSslVaidation: true}
			})

			It("returns a '-k' flag", func() {
				Expect(resultingArgs).To(Equal("-k"))
			})
		})

		Context("when the source doesn't have either CA Cert or SkipSslValidation", func() {
			BeforeEach(func() {
				source = models.Source{}
			})

			It("returns an empty string", func() {
				Expect(resultingArgs).To(Equal(""))
			})
		})
	})

})
