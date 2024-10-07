package store

import (
	"net"
)

type UDPAddressStore = map[string]*net.UDPAddr

var clients = make(UDPAddressStore)

func AddClient(id string, clientAddr *net.UDPAddr) {
	clients[id] = clientAddr
}

func GetClient(id string) *net.UDPAddr {
	return clients[id]
}

func RemoveClient(id string) {
	delete(clients, id)
}

func ClientCount() int {
	return len(clients)
}

func ClientIds() []string {
	ids := make([]string, 0, len(clients))
	for id := range clients {
		ids = append(ids, id)
	}
	return ids
}

func ClientAddresses() []*net.UDPAddr {
	addresses := make([]*net.UDPAddr, 0, len(clients))
	for _, addr := range clients {
		addresses = append(addresses, addr)
	}
	return addresses
}
