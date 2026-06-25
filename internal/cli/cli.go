// Package cli wires command-line arguments to the calc package. It owns all
// input parsing and output formatting (both human-readable and JSON). The calc
// package stays pure; everything user-facing lives here.
package cli

import (
	"fmt"
	"os"
)

// Run is the real entry point. It returns a process exit code so the behaviour
// is scriptable, exactly as the spec requires:
//
//	0 = success / match found
//	1 = no match (used by `check`)
//	2 = usage error or bad input
func Run(args []string) int {
	if len(args) == 0 {
		usage()
		return 2
	}

	switch args[0] {
	case "info":
		return runInfo(args[1:])
	case "split":
		return runSplit(args[1:])
	case "check":
		return runCheck(args[1:])
	case "-h", "--help", "help":
		usage()
		return 0
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %q\n\n", args[0])
		usage()
		return 2
	}
}

func usage() {
	fmt.Fprint(os.Stderr, `goipcalc — IPv4 subnet calculator

usage:
  goipcalc info  <CIDR>           [--json]
  goipcalc split <CIDR> --into N  [--json]
  goipcalc check <A> <B>          [--json]

examples:
  goipcalc info 192.168.1.10/24
  goipcalc split 192.168.1.0/24 --into 26
  goipcalc check 10.0.1.0/24 10.0.1.128/25
  goipcalc check 192.168.1.50 192.168.1.0/24
`)
}
