// Package calc contains the IPv4 subnet arithmetic for goipcalc.
// Everything below is a stub returning errNotImplemented so the whole project
// compiles and runs from day one. Your job is to fill in the function bodies.
// The tests in calc_test.go will go from red to green as you do.
package calc

import (
	"errors"
	"net/netip"
)

// errNotImplemented is returned by every stub. Delete the returns as you
// implement each function.
var errNotImplemented = errors.New("not implemented yet")

// SubnetInfo holds every parameter the `info` command displays.
// The `json:"..."` struct tags control field names when the output is
// serialized for `goipcalc info ... --json`. netip.Addr marshals itself to a
// quoted string automatically, so the JSON comes out clean.
type SubnetInfo struct {
	Address     netip.Addr `json:"address"`
	Network     netip.Addr `json:"network"`
	Prefix      int        `json:"prefix"`
	Netmask     netip.Addr `json:"netmask"`
	Wildcard    netip.Addr `json:"wildcard"`
	Broadcast   netip.Addr `json:"broadcast"`
	HostMin     netip.Addr `json:"host_min"`
	HostMax     netip.Addr `json:"host_max"`
	HostsUsable uint64     `json:"hosts_usable"`
	HostsTotal  uint64     `json:"hosts_total"`
	IsPrivate   bool       `json:"is_private"`
	Class       string     `json:"class"`
}

// Info parses a CIDR like "192.168.1.10/24" and computes all subnet
// parameters. The host part does NOT need to be aligned to the network
// boundary — "192.168.1.77/24" is valid input and Network must come out as
// 192.168.1.0.
func Info(cidr string) (SubnetInfo, error) {
	// HINTS:
	//  1. p, err := netip.ParsePrefix(cidr)  -> gives you p.Addr() and p.Bits().
	//  2. Network address: p.Masked().Addr() zeroes out the host bits for you.
	//  3. To compute broadcast/wildcard you need the raw bytes:
	//        b := p.Addr().As4()   // [4]byte
	//     Build a 4-byte mask from the prefix length, then:
	//        wildcard[i] = ^mask[i]
	//        broadcast[i] = network[i] | ^mask[i]
	//     Turn bytes back into an address with netip.AddrFrom4(...).
	//  4. Counts: total = 1 << (32 - bits); usable = total - 2.
	//     Watch the edge cases: /31 (2 usable, RFC 3021) and /32 (1 address).
	//  5. p.Addr().IsPrivate() already tells you RFC 1918 status.
	//  6. Class: look at the first octet (b[0]):
	//        0–127 = A, 128–191 = B, 192–223 = C, 224–239 = D, 240+ = E.
	return SubnetInfo{}, errNotImplemented
}

// Split divides a network into equal subnets of newPrefix length.
// Example: Split("192.168.1.0/24", 26) -> four /26 networks.
// Must error if newPrefix is not strictly larger than the original prefix,
// or if it is > 32.
func Split(cidr string, newPrefix int) ([]netip.Prefix, error) {
	// HINTS:
	//  1. Parse, then validate: original := p.Bits(); newPrefix must satisfy
	//     original < newPrefix <= 32 — otherwise return a clear error.
	//  2. How many subnets: count := 1 << (newPrefix - original).
	//  3. Step (number of addresses per subnet): step := 1 << (32 - newPrefix).
	//  4. Start at the masked network address. To advance an address by `step`:
	//     convert As4() -> uint32 (binary.BigEndian.Uint32), add step, convert
	//     back with binary.BigEndian.PutUint32 + netip.AddrFrom4.
	//  5. Each subnet is netip.PrefixFrom(addr, newPrefix). Append to a slice.
	return nil, errNotImplemented
}

// Overlaps reports whether networks a and b share ANY address.
func Overlaps(a, b string) (bool, error) {
	// HINT: parse both as netip.Prefix and Masked() them. Two networks overlap
	// iff one contains the other's network address:
	//   pa.Contains(pb.Addr()) || pb.Contains(pa.Addr())
	return false, errNotImplemented
}

// Contains reports whether a single address falls within a network.
func Contains(network, addr string) (bool, error) {
	p, err := netip.ParsePrefix(network)
	if err != nil {
		return false, err
	}
	a, err := netip.ParseAddr(addr)
	if err != nil {
		return false, err
	}
	return p.Contains(a), nil
}
