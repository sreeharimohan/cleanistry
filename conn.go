package main

import (
	"log"
	"net"
)

// AbleToConnect returns a flag indicating whether or not an address is accessible
func AbleToConnect(network string, address string) bool {
	log.Println("Trying to connect to " + address)
	conn, err := net.Dial(network, address)
	log.Println(conn, err)
	if err != nil {
		log.Println("Connection test failed")
		return false
	}
	log.Println("Closing connection")
	conn.Close()
	log.Println("Connection test success")
	return true
}
