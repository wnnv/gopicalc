package calc

import "testing"

// This is the table-driven test pattern — the idiomatic way to test in Go.
// You define a slice of cases, then loop over them. t.Run gives each case its
// own named sub-test so failures point you at the exact row.
//
// Right now Contains is a stub, so this test FAILS (red). Implement Contains
// in calc.go and it goes green. That red->green loop is the whole point.
func TestContains(t *testing.T) {
	tests := []struct {
		name    string
		network string
		addr    string
		want    bool
	}{
		{"addr inside /24", "192.168.1.0/24", "192.168.1.50", true},
		{"addr outside /24", "192.168.1.0/24", "192.168.2.50", false},
		{"upper boundary /8", "10.0.0.0/8", "10.255.255.255", true},
		{"just outside /8", "10.0.0.0/8", "11.0.0.1", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Contains(tt.network, tt.addr)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Contains(%q, %q) = %v, want %v",
					tt.network, tt.addr, got, tt.want)
			}
		})
	}
}

// TODO: write the same style of table for Info, Split and Overlaps.
// For Info, assert individual fields, e.g.:
//   {"basic /24", "192.168.1.10/24", "192.168.1.0", "192.168.1.255", 254}
// and check out.Network.String(), out.Broadcast.String(), out.HostsUsable.
// Make sure to cover the edge cases: /31, /32, /0, and an unaligned host
// like "192.168.1.77/24".
