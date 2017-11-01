package main

import (
	"net"
)

// AbleToConnect returns a flag indicating whether or not an address is accessible
func AbleToConnect(network string, address string) bool {
	conn, err := net.Dial(network, address)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
