# OVH Resource Provider

The OVH Resource Provider lets you manage [OVHcloud](https://www.ovhcloud.com/en/) resources.

<a href="https://github.com/ovh/pulumi-ovh/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/v/release/ovh/pulumi-ovh?logo=github&style=flat-square"></a>
[![GoDoc](https://godoc.org/github.com/ovh/pulumi-ovh?status.svg)](https://godoc.org/github.com/ovh/pulumi-ovh)
[![NPM version](https://badge.fury.io/js/@scraly%2Fpulumi-ovh.svg)](https://badge.fury.io/js/@ovh%2Fpulumi-ovh)
[![PyPI version](https://badge.fury.io/py/pulumi-ovh.svg)](https://badge.fury.io/py/pulumi-ovh)
[![NuGet version](https://badge.fury.io/nu/Pulumi.Ovh.svg)](https://badge.fury.io/nu/Pulumi.Ovh)
<a href="https://gitpod.io/#https://github.com/ovh/pulumi-ovh"><img src="https://img.shields.io/badge/Contribute%20with-Gitpod-908a85?logo=gitpod" alt="Contribute with Gitpod"/></a>

## Usage

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @scraly/pulumi-ovh
```

or `yarn`:

```bash
yarn add @scraly/pulumi-ovh
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi-ovh
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/ovh/pulumi-ovh/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Ovh
```

## Configuration

The following configuration points are available for the `Ovh` provider:

- `ovh:endpoint` (environment: `OVH_ENDPOINT`) - the Ovh endpoint, such `ovh-eu`
- `ovh:applicationKey` (environment: `OVH_APPLICATION_KEY`) - the Ovh application key
- `ovh:applicationSecret` (environment: `OVH_APPLICATION_SECRET`) - the Ovh application secret
- `ovh:consumerKey` (environment: `OVH_CONSUMER_KEY`) - the Ovh consumer key

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/ovh/api-docs/).
