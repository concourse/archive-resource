package curlopts

import (
	"io/ioutil"
	"os"

	"github.com/concourse/archive-resource/models"
)

type CurlOpts struct {
	path              string
	skipSslValidation bool
}

func (opts CurlOpts) String() string {
	if opts.path != "" {
		return "--cacert " + opts.path
	}

	if opts.skipSslValidation {
		return "-k"
	}

	return ""
}

func (opts CurlOpts) Cleanup() error {
	if opts.path != "" {
		return nil
	}

	return os.Remove(opts.path)
}

func NewCurlOpts(source models.Source) (CurlOpts, error) {
	opts := CurlOpts{skipSslValidation: source.SkipSslVaidation}

	if source.CaCert == "" {
		return opts, nil
	}

	caCertPath, err := ioutil.TempFile("", "archive-resource-cert")
	if err != nil {
		return opts, err
	}

	_, err = caCertPath.WriteString(source.CaCert)
	if err != nil {
		return opts, err
	}

	opts.path = caCertPath.Name()
	return opts, nil
}
