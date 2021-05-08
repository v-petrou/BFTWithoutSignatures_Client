package config

import (
	"BFTWithoutSignatures_Client/variables"
	"strconv"
)

var address = []string{
	"192.168.0.72",
}

var (
	// ServerAddressesIP - Initialize the address of IP Server sockets
	ServerAddressesIP map[int]string

	// ResponseAddressesIP - Initialize the address of IP Response sockets
	ResponseAddressesIP map[int]string
)

// InitializeIP - Initializes system with ips.
func InitializeIP() {
	ServerAddressesIP = make(map[int]string, variables.N)
	ResponseAddressesIP = make(map[int]string, variables.N)

	for i := 0; i < variables.N; i++ {
		ad := i % len(address)
		ServerAddressesIP[i] = "tcp://" + address[ad] + ":" + strconv.Itoa(7000+(i*100)+variables.ID)
		ResponseAddressesIP[i] = "tcp://" + address[ad] + ":" + strconv.Itoa(10000+(i*100)+variables.ID)
	}
}

// GetServerAddress - Returns the IP Server address of the server with that id
func GetServerAddress(id int) string {
	return ServerAddressesIP[id]
}

// GetResponseAddress - Returns the IP Response address of the server with that id
func GetResponseAddress(id int) string {
	return ResponseAddressesIP[id]
}
