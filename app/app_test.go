package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewNetworkInfos uses table driven to unit testing NewNetworkInfos function.
func TestNewNetworkInfos(t *testing.T) {
	cases := []struct {
		input             string
		expectedAddr      string
		expectedPrefixLen int
		expectErr         bool
	}{
		{
			input:             "10.0.2.0/24",
			expectedAddr:      "10.0.2.0",
			expectedPrefixLen: 24,
			expectErr:         false,
		},
		{
			input:             "10.0.2.10/20",
			expectedAddr:      "10.0.0.0",
			expectedPrefixLen: 20,
			expectErr:         false,
		},
		{
			input:     "10.0.2.0/",
			expectErr: true,
		},
		{
			input:     "10.0.2.0",
			expectErr: true,
		},
		{
			input:     "10.0.2.",
			expectErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := NewNetworkInfos(tc.input)
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedAddr, got.Address.String())
				assert.Equal(t, tc.expectedPrefixLen, got.PrefixLength)
			}
		})
	}
}

// TestIsSameAs uses table driven to unit testing IsSameAs function.
func TestIsSameAs(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect bool
	}{
		{
			name:   "same ipv4 prefixes",
			input:  []string{"10.0.2.0/24", "10.0.2.10/24"},
			expect: true,
		},
		{
			name:   "not same ipv4 prefixes",
			input:  []string{"192.168.1.0/24", "10.10.10.0/24"},
			expect: false,
		},
		{
			name:   "same ipv6 prefixes",
			input:  []string{"fe80::c845:ea23:ad2e/64", "fe80::c845:ea23:ad2e:65b8/64"},
			expect: true,
		},
		{
			name:   "not same ipv6 prefixes",
			input:  []string{"fe80::c845:ea23:ad2e/128", "fe80::c845:ea23:ad2e:65b8/128"},
			expect: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			n1, err := NewNetworkInfos(tc.input[0])
			assert.NoError(t, err)
			n2, err := NewNetworkInfos(tc.input[1])
			assert.NoError(t, err)
			got := n2.IsSameAs(n1)
			assert.Equal(t, tc.expect, got)
		})
	}
}

// TestIsSubsetOf uses table driven to unit testing IsSubsetOf function.
func TestIsSubsetOf(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect bool
	}{
		{
			name:   "subset ipv4 prefixes",
			input:  []string{"10.0.0.0/8", "10.0.2.10/24"},
			expect: true,
		},
		{
			name:   "not subset ipv4 prefixes",
			input:  []string{"192.168.1.0/24", "192.168.1.0/23"},
			expect: false,
		},
		{
			name:   "subset ipv6 prefixes",
			input:  []string{"fe80::c845:ea23:ad2e/64", "fe80::c845:ea23:ad2e:65b8/128"},
			expect: true,
		},
		{
			name:   "not subset ipv6 prefixes",
			input:  []string{"fe80::c845:ea23:ad2e/128", "fe80::c845:ea23:ad2e:65b8/64"},
			expect: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			n1, err := NewNetworkInfos(tc.input[0])
			assert.NoError(t, err)
			n2, err := NewNetworkInfos(tc.input[1])
			assert.NoError(t, err)
			got := n2.IsSubsetOf(n1)
			assert.Equal(t, tc.expect, got)
		})
	}
}

// TestIsSupersetOf uses table driven to unit testing IsSupersetOf function.
func TestIsSupersetOf(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect bool
	}{
		{
			name:   "superset ipv4 prefixes",
			input:  []string{"10.10.10.0/24", "10.0.2.10/8"},
			expect: true,
		},
		{
			name:   "not superset ipv4 prefixes",
			input:  []string{"0.0.0.0/0", "10.0.0.0/8"},
			expect: false,
		},
		{
			name:   "superset ipv6 prefixes",
			input:  []string{"fe80::c845:ea23:ad2e/64", "fe80::/8"},
			expect: true,
		},
		{
			name:   "not superset ipv6 prefixes",
			input:  []string{"::/0", "::1/128"},
			expect: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			n1, err := NewNetworkInfos(tc.input[0])
			assert.NoError(t, err)
			n2, err := NewNetworkInfos(tc.input[1])
			assert.NoError(t, err)
			got := n2.IsSupersetOf(n1)
			assert.Equal(t, tc.expect, got)
		})
	}
}

// TestCheckOverlapStatus uses table driven to unit testing CheckOverlapStatus function.
func TestCheckOverlapStatus(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect OverlapStatus
	}{
		{
			name:   "same ipv4 prefixes",
			input:  []string{"10.0.2.0/24", "10.0.2.10/24"},
			expect: SAME,
		},
		{
			name:   "subset ipv4 prefixes",
			input:  []string{"10.0.0.0/8", "10.0.2.10/24"},
			expect: SUBSET,
		},
		{
			name:   "superset ipv4 prefixes",
			input:  []string{"10.10.10.0/24", "10.0.2.10/8"},
			expect: SUPERSET,
		},
		{
			name:   "different ipv4 prefixes",
			input:  []string{"172.16.0.0/24", "192.168.0.0/24"},
			expect: DIFFERENT,
		},
		{
			name:   "same ipv6 prefixes",
			input:  []string{"fe80::c845:ea23:ad2e/64", "fe80::c845:ea23:ad2e:65b8/64"},
			expect: SAME,
		},
		{
			name:   "subset ipv6 prefixes",
			input:  []string{"fe80::c845:ea23:ad2e/64", "fe80::c845:ea23:ad2e:65b8/128"},
			expect: SUBSET,
		},
		{
			name:   "superset ipv6 prefixes",
			input:  []string{"fe80::c845:ea23:ad2e/64", "fe80::/8"},
			expect: SUPERSET,
		},
		{
			name:   "same ipv6 prefixes",
			input:  []string{"fe80::/8", "::1/128"},
			expect: DIFFERENT,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			n1, err := NewNetworkInfos(tc.input[0])
			assert.NoError(t, err)
			n2, err := NewNetworkInfos(tc.input[1])
			assert.NoError(t, err)
			got := n2.CheckOverlapStatus(n1)
			assert.Equal(t, tc.expect, got)
		})
	}
}

// TestIsComparableTo uses table driven to unit testing IsComparableTo function.
func TestIsComparableTo(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect bool
	}{
		{
			name:   "should compare v4tov4",
			input:  []string{"10.0.2.0/24", "192.168.0.0/0"},
			expect: true,
		},
		{
			name:   "should compare v6tov6",
			input:  []string{"fd74:5909:aafe::35f0:65ca:6fa1/128", "fe80::c845:ea23:ad2e:65b8/64"},
			expect: true,
		},
		{
			name:   "should not compare v4tov6",
			input:  []string{"10.0.2.0/24", "::/0"},
			expect: false,
		},
		{
			name:   "should not compare v6tov4",
			input:  []string{"::/0", "0.0.0.0/0"},
			expect: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			n1, err := NewNetworkInfos(tc.input[0])
			assert.NoError(t, err)
			n2, err := NewNetworkInfos(tc.input[1])
			assert.NoError(t, err)
			got := n2.IsComparableTo(n1)
			assert.Equal(t, tc.expect, got)
		})
	}
}

// TestInit uses table driven to unit testing Init function.
func TestInit(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect []string
	}{
		{
			name:   "empty build flag",
			input:  []string{"", "", ""},
			expect: []string{"<undefined>", "<undefined>", "<undefined>"},
		},
		{
			name:   "non empty build flag",
			input:  []string{"2022-07-17 11:58:04 AM GMT", "c48f653", "v1.0"},
			expect: []string{"2022-07-17 11:58:04 AM GMT", "c48f653", "v1.0"},
		},
		{
			name:   "mixed values build flag",
			input:  []string{"", "c48f653", "v1.0"},
			expect: []string{"<undefined>", "c48f653", "v1.0"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for i := range tc.input {
				Init(&tc.input[i])
				assert.Equal(t, tc.expect[i], tc.input[i])
			}
		})
	}
}

// TestRun uses table driven to unit testing Run function.
func TestRun(t *testing.T) {
	cases := []struct {
		name           string
		input          []string
		expectedStatus OverlapStatus
		expectErr      bool
	}{
		{
			name:           "same ipv4 prefixes",
			input:          []string{"10.0.2.0/24", "10.0.2.10/24"},
			expectedStatus: SAME,
			expectErr:      false,
		},
		{
			name:           "subset ipv4 prefixes",
			input:          []string{"10.0.0.0/8", "10.0.2.10/24"},
			expectedStatus: SUBSET,
			expectErr:      false,
		},
		{
			name:           "superset ipv4 prefixes",
			input:          []string{"10.10.10.0/24", "10.0.2.10/8"},
			expectedStatus: SUPERSET,
			expectErr:      false,
		},
		{
			name:           "different ipv4 prefixes",
			input:          []string{"172.16.0.0/24", "192.168.0.0/24"},
			expectedStatus: DIFFERENT,
			expectErr:      false,
		},
		{
			name:           "same ipv6 prefixes",
			input:          []string{"fe80::c845:ea23:ad2e/64", "fe80::c845:ea23:ad2e:65b8/64"},
			expectedStatus: SAME,
			expectErr:      false,
		},
		{
			name:           "subset ipv6 prefixes",
			input:          []string{"fe80::c845:ea23:ad2e/64", "fe80::c845:ea23:ad2e:65b8/128"},
			expectedStatus: SUBSET,
			expectErr:      false,
		},
		{
			name:           "superset ipv6 prefixes",
			input:          []string{"fe80::c845:ea23:ad2e/64", "fe80::/8"},
			expectedStatus: SUPERSET,
			expectErr:      false,
		},
		{
			name:           "same ipv6 prefixes",
			input:          []string{"fe80::/8", "::1/128"},
			expectedStatus: DIFFERENT,
			expectErr:      false,
		},
		{
			name:           "not comparable prefixes",
			input:          []string{"::/0", "0.0.0.0/0"},
			expectedStatus: NOTFOUND,
			expectErr:      true,
		},
		{
			name:           "invalid first prefix",
			input:          []string{"10.10.0.0", "0.0.0.0/0"},
			expectedStatus: NOTFOUND,
			expectErr:      true,
		},
		{
			name:           "invalid second prefix",
			input:          []string{"192.168.0.0/24", "10.10.0.0/"},
			expectedStatus: NOTFOUND,
			expectErr:      true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			status, err := Run(tc.input[0], tc.input[1])
			if tc.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedStatus, status)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedStatus, status)
			}
		})
	}
}
