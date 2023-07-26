# Squadcast Resource Provider

The Squadcast Resource Provider lets you manage [Squadcast](https://www.squadcast.com/) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @valisinsights/pulumi-squadcast
```

or `yarn`:

```bash
yarn add @valisinsights/pulumi-squadcast
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi-squadcast
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/valisinsights/pulumi-squadcast/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package VALISInsights.Pulumi.Squadcast
```

## Configuration

The following configuration points are available for the `squadcast` provider:

- `squadcast:refreshToken` (environment: `SQUADCAST_REFRESH_TOKEN `) - the API key for `squadcast`
- `squadcast:region` (environment: `SQUADCAST_REGION`) - the region in which to deploy resources
