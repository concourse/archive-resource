# Archive Resource

Simply downloads and extracts an archive to the destination.

**NOTE**: This resource is only intended for use in `fly execute` (it's how
your inputs get uploaded). It won't work in a pipeline because `check` never
yields any valid versions. This is because a download URL is not enough to
continuously integrate with something, since the endpoint isn't versioned.
You probably want the [S3 resource](https://github.com/concourse/s3-resource)
or the [GitHub Release
resource](https://github.com/concourse/github-release-resource) instead.

## Source Configuration

* `uri`: *Required.* The location of the file to download.
* `ca_cert`: *Optional.* The contents of server CA cert.
* `skip_ssl_validation`: *Optional.* Skip SSL validation.

## Behavior

### `check`: Not implemented.

As this resource is mainly used for one-off downloads (with
[Fly](https://github.com/concourse/fly)), there aren't really any versioning
semantics.


### `in`: Download and extract the archive.

Fetches a `.tar.gz` file from the URL, and extracts it to the destination as
it's downloading.


#### Parameters

*None.*


### `out`: Not implemented.

Currently there is no output functionality. In principle, this could be
configured with a directory to compress and upload to the `uri`, however
this is not currently implemented.

#### Parameters

*None.*

## Development

### Prerequisites

* golang is *required* - version 1.9.x is tested; earlier versions may also
  work.
* docker is *required* - version 17.06.x is tested; earlier versions may also
  work.
* godep is used for dependency management of the golang packages.

### Running the tests

The tests have been embedded with the `Dockerfile`; ensuring that the testing
environment is consistent across any `docker` enabled platform. When the docker
image builds, the test are run inside the docker container, on failure they
will stop the build.

Run the tests with the following command:

```sh
docker build -t archive-resource .
```

### Contributing

Please make all pull requests to the `master` branch and ensure tests pass
locally.
