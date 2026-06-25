# goipcalc

IPv4 subnet calculator CLI. Computes subnet parameters, splits networks into
smaller subnets, and checks for overlap / containment.

> Pet project for moving from PHP to Go. The `calc` package is the same subnet
> math that later powers `ipam-api`.

## Build & run

```bash
go run ./cmd/goipcalc info 192.168.1.10/24
go build -o goipcalc ./cmd/goipcalc   # produces ./goipcalc
```

## Commands

```bash
# subnet parameters
goipcalc info 192.168.1.10/24
goipcalc info 192.168.1.10/24 --json

# split a network into smaller subnets
goipcalc split 192.168.1.0/24 --into 26

# overlap / containment
goipcalc check 10.0.1.0/24 10.0.1.128/25     # two networks -> overlap?
goipcalc check 192.168.1.50 192.168.1.0/24   # address in network?
```

Exit codes: `0` success/match, `1` no match (for `check`), `2` bad input.

## Develop

```bash
go test ./...      # run tests
go vet ./...       # static checks
go mod tidy        # tidy dependencies
```

## Status

The CLI wiring is complete. The arithmetic in `internal/calc` is stubbed —
fill in `Info`, `Split`, `Overlaps`, `Contains` (hints are in the code) and
watch the tests go green.
