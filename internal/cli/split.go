package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/wnnv/goipcalc/internal/calc"
)

func runSplit(args []string) int {
	fs := flag.NewFlagSet("split", flag.ContinueOnError)
	into := fs.Int("into", 0, "target prefix length, e.g. 26")
	asJSON := fs.Bool("json", false, "output as JSON")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if fs.NArg() != 1 || *into == 0 {
		fmt.Fprintln(os.Stderr, "split: usage: goipcalc split <CIDR> --into N")
		return 2
	}

	subnets, err := calc.Split(fs.Arg(0), *into)
	if err != nil {
		fmt.Fprintf(os.Stderr, "split: %v\n", err)
		return 2
	}

	if *asJSON {
		out := make([]string, len(subnets))
		for i, s := range subnets {
			out[i] = s.String()
		}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		_ = enc.Encode(out)
		return 0
	}

	for _, s := range subnets {
		// Once calc.Info works you can enrich each line with host count and
		// range (call calc.Info(s.String()) and format like the spec shows).
		// For now just print the subnet prefix.
		fmt.Println(s.String())
	}
	return 0
}
