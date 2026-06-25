package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/wnnv/goipcalc/internal/calc"
)

func runCheck(args []string) int {
	fs := flag.NewFlagSet("check", flag.ContinueOnError)
	asJSON := fs.Bool("json", false, "output as JSON")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if fs.NArg() != 2 {
		fmt.Fprintln(os.Stderr, "check: need exactly two arguments")
		return 2
	}

	a, b := fs.Arg(0), fs.Arg(1)

	// If exactly one argument is a bare address (no "/"), treat it as a
	// "does this address belong to this network" check. Otherwise compare two
	// networks for overlap.
	switch {
	case !strings.Contains(a, "/") && strings.Contains(b, "/"):
		return reportContains(b, a, *asJSON)
	case strings.Contains(a, "/") && !strings.Contains(b, "/"):
		return reportContains(a, b, *asJSON)
	default:
		return reportOverlap(a, b, *asJSON)
	}
}

func reportOverlap(a, b string, asJSON bool) int {
	overlap, err := calc.Overlaps(a, b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "check: %v\n", err)
		return 2
	}
	if asJSON {
		_ = json.NewEncoder(os.Stdout).Encode(map[string]bool{"overlap": overlap})
	} else if overlap {
		fmt.Println("Overlap: YES")
	} else {
		fmt.Println("Overlap: NO")
	}
	if overlap {
		return 0
	}
	return 1
}

func reportContains(network, addr string, asJSON bool) int {
	contains, err := calc.Contains(network, addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "check: %v\n", err)
		return 2
	}
	if asJSON {
		_ = json.NewEncoder(os.Stdout).Encode(map[string]bool{"contains": contains})
	} else if contains {
		fmt.Println("Contains: YES")
	} else {
		fmt.Println("Contains: NO")
	}
	if contains {
		return 0
	}
	return 1
}
