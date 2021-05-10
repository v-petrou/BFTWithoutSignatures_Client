package config

import (
	"BFTWithoutSignatures_Client/variables"
	"strconv"
)

var address = []string{}

var (
	// ServerAddressesIP - Initialize the address of IP Server sockets
	ServerAddressesIP map[int]string

	// ResponseAddressesIP - Initialize the address of IP Response sockets
	ResponseAddressesIP map[int]string
)

// InitializeIP - Initializes system with ips.
func InitializeIP(clients int) {
	ServerAddressesIP = make(map[int]string, variables.N)
	ResponseAddressesIP = make(map[int]string, variables.N)

	for i := 0; i < variables.N; i++ {
		ad := i % len(address)
		ServerAddressesIP[i] = "tcp://" + address[ad] + ":" + strconv.Itoa(27015+variables.ID+(i*clients))
		ResponseAddressesIP[i] = "tcp://" + address[ad] + ":" + strconv.Itoa(27065+variables.ID+(i*clients))
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
