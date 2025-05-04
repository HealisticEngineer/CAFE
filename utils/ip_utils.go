package utils

import (
	"net"
)

// Allowed CIDRs for incoming client IPs
var allowedCIDRs = []string{
	"192.168.0.0/16",
	"10.0.0.0/8",
	"172.16.0.0/12",
}

// IPAllowed checks if the given IP is allowed
func IPAllowed(remoteAddr string) bool {
	ipStr, _, _ := net.SplitHostPort(remoteAddr)
	ip := net.ParseIP(ipStr)
	for _, cidr := range allowedCIDRs {
		_, network, _ := net.ParseCIDR(cidr)
		if network.Contains(ip) {
			return true
		}
	}
	return false
}
