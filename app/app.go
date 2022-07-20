// Package app provides custom types and methods to evaluate the overlap
// status between two prefixes from same types (v4 or v6).
package app

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// OverlapStatus is custom type to restrict possible status values.
type OverlapStatus string

// Expected overlap statuses.
const (
	SUBSET    OverlapStatus = "subset"
	SUPERSET  OverlapStatus = "superset"
	SAME      OverlapStatus = "same"
	DIFFERENT OverlapStatus = "different"
	NOTFOUND  OverlapStatus = ""
)

// NetworkInfos holds customized network details.
type NetworkInfos struct {
	CIDR         *net.IPNet // Formatted IP Prefix.
	Address      net.IP     // Real Network IP address.
	PrefixLength int        // Prefix length associated.
}

// NewNetworkInfos returns an instance of NetworkInfos object.
func NewNetworkInfos(input string) (*NetworkInfos, error) {
	_, cidr, err := net.ParseCIDR(input)
	if err != nil {
		return &NetworkInfos{}, err
	}
	n := &NetworkInfos{
		CIDR:    cidr,
		Address: cidr.IP,
	}
	prefixLength, err := strconv.Atoi(strings.Split(input, "/")[1])
	n.PrefixLength = prefixLength
	return n, err
}

// IsSame method verifies if its network equal a given network.
func (n2 *NetworkInfos) IsSameAs(n1 *NetworkInfos) bool {
	if n2.Address.Equal(n1.Address) && n1.PrefixLength == n2.PrefixLength {
		return true
	}
	return false
}

// IsSubset method verifies if its network is included into a given network.
func (n2 *NetworkInfos) IsSubsetOf(n1 *NetworkInfos) bool {
	if n1.CIDR.Contains(n2.Address) && n1.PrefixLength <= n2.PrefixLength {
		return true
	}
	return false
}

// IsSuperset method verifies if its network includes a given network.
func (n2 *NetworkInfos) IsSupersetOf(n1 *NetworkInfos) bool {
	if n2.CIDR.Contains(n1.Address) && n2.PrefixLength <= n1.PrefixLength {
		return true
	}
	return false
}

// CheckOverlapStatus method calls subsequently others methods to returns the overlap status.
func (n2 *NetworkInfos) CheckOverlapStatus(n1 *NetworkInfos) OverlapStatus {
	switch {
	case n2.IsSameAs(n1):
		return SAME
	case n2.IsSubsetOf(n1):
		return SUBSET
	case n2.IsSupersetOf(n1):
		return SUPERSET
	default:
		return DIFFERENT
	}
}

// IsComparableTo function verifies two network prefixes are from same v4 or v6 type.
func (n2 *NetworkInfos) IsComparableTo(n1 *NetworkInfos) bool {
	return strings.Contains(n2.Address.String(), ":") == strings.Contains(n1.Address.String(), ":")
}

// Init function asserts and sets default value of build flags.
func Init(flags ...*string) {
	for _, f := range flags {
		if *f == "" {
			*f = "<undefined>"
		}
	}
}

// Run processes user-provided networks and returns the overlap status.
func Run(first string, second string) (OverlapStatus, error) {
	n1, err := NewNetworkInfos(first)
	if err != nil {
		return NOTFOUND, fmt.Errorf("failed to construct network infos: %v", err)
	}
	n2, err := NewNetworkInfos(second)
	if err != nil {
		return NOTFOUND, fmt.Errorf("failed to construct network infos: %v", err)
	}

	if !n2.IsComparableTo(n1) {
		return NOTFOUND, fmt.Errorf("failed to evaluate: prefixes are not from same type (v4 or v6)")
	}
	return n2.CheckOverlapStatus(n1), nil
}
