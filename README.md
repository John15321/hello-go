# hello-go

Prints out Hello World! A small Go CLI used as a testing ground for a Go project setup.

## Quick start

```bash
# Fetch deps + dev tools
make dev-install

# Run
go run ./cmd/hello-go
go run ./cmd/hello-go --name Jan

# Build
make build
./bin/hello-go --help
```

## Development

Everything is driven by the `Makefile`. Common targets:

| Target             | What it does                                         |
| ------------------ | ---------------------------------------------------- |
| `make install`     | Download and verify module dependencies              |
| `make dev-install` | Sync deps and dev tool directives (`go mod tidy`)    |
| `make test`        | Run tests with coverage (`gotestsum` + race)         |
| `make test-quick`  | Run tests without coverage                           |
| `make lint`        | `golangci-lint run`                                  |
| `make lint-fix`    | `golangci-lint run --fix`                            |
| `make format`      | `gofmt` + `goimports` via `golangci-lint fmt`        |
| `make format-check`| Fail if code is not formatted                        |
| `make vet`         | `go vet ./...`                                       |
| `make check`       | format-check + lint + vet + test                     |
| `make fix`         | format + lint-fix                                    |
| `make build`       | Build binary into `./bin/hello-go`                   |
| `make run`         | `go run ./cmd/hello-go` (pass args with `ARGS='...'`)|
| `make release-snapshot` | Local goreleaser snapshot                       |
| `make clean`       | Remove build / test artifacts                        |
| `make all`         | `dev-install` + `fix` + `check` + `build`            |
| `make ci`          | Alias for `check`                                    |

Dev tools (`golangci-lint`, `gotestsum`, `goreleaser`) are declared as
[Go 1.24+ tool dependencies](https://go.dev/blog/tool-dependencies) in
`go.mod`, so `go tool <name>` uses the versions pinned by the module — no
global installs required.

## Layout

```
cmd/hello-go/     entrypoint (main package)
internal/cli/     cobra commands (not importable outside this module)
```

## License

MIT — see [LICENSE](LICENSE).
