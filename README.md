# internal-devbox

Internal dev-env helper CLI for Pantalasa engineers. **Not customer-facing. Not deployed to production.** Runs on engineer laptops only.

## Installation

```bash
go install github.com/pantalasa/internal-devbox@latest
```

Or build from source:

```bash
go mod download
go build -o devbox .
```

## Usage

```bash
devbox env            # Print recommended local env vars
devbox version        # Print tool version
devbox help           # Show help
```

## Contributing

1. Fork / branch
2. `go test ./...`
3. Open a PR

## Governance

This repo is registered as a component in [`pantalasa/lunar`](https://github.com/pantalasa/lunar) under the `engineering.internal` domain. It is tagged `internal`, which exempts it from strict enforcement of the supply-chain and PR-hygiene policies that apply to customer-facing services. Those checks still run here but only contribute to the dashboard score — they never block a PR merge.
