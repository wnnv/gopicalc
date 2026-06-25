package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/wnnv/goipcalc/internal/calc"
)

// Each subcommand gets its OWN flag.FlagSet. This is the standard-library way
// to do subcommands without any third-party dependency. flag.ContinueOnError
// makes Parse return an error instead of calling os.Exit, so we control the
// exit code ourselves.
func runInfo(args []string) int {
	fs := flag.NewFlagSet("info", flag.ContinueOnError)
	asJSON := fs.Bool("json", false, "output as JSON")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if fs.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "info: need exactly one CIDR argument")
		return 2
	}

	out, err := calc.Info(fs.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "info: %v\n", err)
		return 2
	}

	if *asJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		_ = enc.Encode(out)
		return 0
	}

	fmt.Printf("Address:     %s\n", out.Address)
	fmt.Printf("Network:     %s/%d\n", out.Network, out.Prefix)
	fmt.Printf("Netmask:     %s\n", out.Netmask)
	fmt.Printf("Wildcard:    %s\n", out.Wildcard)
	fmt.Printf("Broadcast:   %s\n", out.Broadcast)
	fmt.Printf("HostMin:     %s\n", out.HostMin)
	fmt.Printf("HostMax:     %s\n", out.HostMax)
	fmt.Printf("Hosts:       %d (всего адресов: %d)\n", out.HostsUsable, out.HostsTotal)

	access := "Public"
	if out.IsPrivate {
		access = "Private (RFC 1918)"
	}
	fmt.Printf("Type:        %s, Class %s\n", access, out.Class)
	return 0
}
