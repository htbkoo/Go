package main

import (
	"fmt"
)

type IPAddr [4]byte

// TODO: Add a "String" method that is defined on values of type "IPAddr" (with "ip" being the receiver) and returns a string.

/*
func (ip IPAddr) String() string {
	var sb bytes.Buffer
	sb.WriteString(fmt.Sprintf("%d", ip[0]))

	for i, p := range ip {
		if i > 0 {
			sb.WriteString(fmt.Sprintf(".%d", p))
		}
	}

	return sb.String()
}
*/

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
