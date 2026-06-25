package main

import (
	"os"

	"github.com/wnnv/goipcalc/internal/cli"
)

// main is intentionally tiny: it does NOTHING except hand the args to the cli
// package and use whatever exit code comes back as the process exit code.
// All real logic lives in internal/. This is the idiomatic Go layout.
func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
