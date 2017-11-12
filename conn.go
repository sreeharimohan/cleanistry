package main

import (
	"log"
	"net"
)

// AbleToConnect returns a flag indicating whether or not an address is accessible
func AbleToConnect(network string, address string) bool {
	log.Println(address)
	conn, err := net.Dial(network, address)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
