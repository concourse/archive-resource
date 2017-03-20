package curlopts

import (
	"io/ioutil"
	"os"

	"github.com/concourse/archive-resource/models"
)

type CurlOpts struct {
	path string
}

func (opts CurlOpts) String() string {
	if opts.path == "" {
		return ""
	}

	return "--cacert " + opts.path
}

func (opts CurlOpts) Cleanup() error {
	if opts.path != "" {
		return nil
	}

	return os.Remove(opts.path)
}

func NewCurlOpts(source models.Source) (CurlOpts, error) {
	if source.CaCert == "" {
		return CurlOpts{}, nil
	}

	caCertPath, err := ioutil.TempFile("", "archive-resource-cert")
	if err != nil {
		return CurlOpts{}, err
	}

	_, err = caCertPath.WriteString(source.CaCert)
	if err != nil {
		return CurlOpts{}, err
	}

	return CurlOpts{path: caCertPath.Name()}, nil
}
