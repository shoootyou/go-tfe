HCP Terraform and Terraform Enterprise Go Client
==============================

[![Tests](https://github.com/shoootyou/go-tfe/actions/workflows/ci.yml/badge.svg)](https://github.com/shoootyou/go-tfe/actions/workflows/ci.yml)
[![GitHub license](https://img.shields.io/github/license/shoootyou/go-tfe.svg)](https://github.com/shoootyou/go-tfe/blob/main/LICENSE)
[![GoDoc](https://godoc.org/github.com/shoootyou/go-tfe?status.svg)](https://godoc.org/github.com/shoootyou/go-tfe)
[![Go Report Card](https://goreportcard.com/badge/github.com/shoootyou/go-tfe)](https://goreportcard.com/report/github.com/shoootyou/go-tfe)
[![GitHub issues](https://img.shields.io/github/issues/shoootyou/go-tfe.svg)](https://github.com/shoootyou/go-tfe/issues)

The official Go API client for [HCP Terraform and Terraform Enterprise](https://www.hashicorp.com/products/terraform).

This client supports the [HCP Terraform V2 API](https://developer.hashicorp.com/terraform/cloud-docs/api-docs).
As Terraform Enterprise is a self-hosted distribution of HCP Terraform, this
client supports both HCP Terraform and Terraform Enterprise use cases. In all package
documentation and API, the platform will always be stated as 'Terraform
Enterprise' - but a feature will be explicitly noted as only supported in one or
the other, if applicable (rare).

## Version Information

Almost always, minor version changes will indicate backwards-compatible features and enhancements. Occasionally, function signature changes that reflect a bug fix may appear as a minor version change. Patch version changes will be used for bug fixes, performance improvements, and otherwise unimpactful changes.

## Example Usage

Construct a new TFE client, then use the various endpoints on the client to
access different parts of the Terraform Enterprise API. The following example lists
all organizations.

### (Recommended Approach) Using custom config to provide configuration details to the API client

```go
import (
  "context"
  "log"

  "github.com/shoootyou/go-tfe"
)

config := &tfe.Config{
	Address: "https://tfe.local",
	Token: "insert-your-token-here",
  RetryServerErrors: true,
}

client, err := tfe.NewClient(config)
if err != nil {
	log.Fatal(err)
}

orgs, err := client.Organizations.List(context.Background(), nil)
if err != nil {
	log.Fatal(err)
}
```

### Using the default config with env vars
The default configuration makes use of the `TFE_ADDRESS` and `TFE_TOKEN` environment variables.

1. `TFE_ADDRESS` - URL of a HCP Terraform or Terraform Enterprise instance. Example: `https://tfe.local`
1. `TFE_TOKEN` - An [API token](https://developer.hashicorp.com/terraform/cloud-docs/users-teams-organizations/api-tokens) for the HCP Terraform or Terraform Enterprise instance.

**Note:** Alternatively, you can set `TFE_HOSTNAME` which serves as a fallback for `TFE_ADDRESS`. It will only be used if `TFE_ADDRESS` is not set and will resolve the host to an `https` scheme. Example: `tfe.local` => resolves to `https://tfe.local`

The environment variables are used as a fallback to configure TFE client if the Address or Token values are not provided as in the cases below:

#### Using the default configuration
```go
import (
  "context"
  "log"

  "github.com/shoootyou/go-tfe"
)

// Passing nil to tfe.NewClient method will also use the default configuration
client, err := tfe.NewClient(tfe.DefaultConfig())
if err != nil {
	log.Fatal(err)
}

orgs, err := client.Organizations.List(context.Background(), nil)
if err != nil {
	log.Fatal(err)
}
```

#### When Address or Token has no value
```go
import (
  "context"
  "log"

  "github.com/shoootyou/go-tfe"
)

config := &tfe.Config{
	Address: "",
	Token: "",
}

client, err := tfe.NewClient(config)
if err != nil {
	log.Fatal(err)
}

orgs, err := client.Organizations.List(context.Background(), nil)
if err != nil {
	log.Fatal(err)
}
```

## Documentation

For complete usage of the API client, see the [full package docs](https://pkg.go.dev/github.com/shoootyou/go-tfe).

## Examples

See the [examples directory](https://github.com/shoootyou/go-tfe/tree/main/examples).

## Running tests

See [TESTS.md](docs/TESTS.md).

## Issues and Contributing

See [CONTRIBUTING.md](docs/CONTRIBUTING.md)

## Releases

See [RELEASES.md](docs/RELEASES.md)
