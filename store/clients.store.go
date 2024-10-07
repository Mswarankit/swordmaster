package store

import (
	"net"
	"swordmaster/models"
)

type UDPAddressStore = map[string]*models.Client

var clients = make(UDPAddressStore)

func AddClient(id string, clientAddr *net.UDPAddr) {
	clients[id] = models.NewClient(id, clientAddr)
}

func GetClient(id string) *models.Client {
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

func GetClients() []*models.Client {
	output := make([]*models.Client, 0, len(clients))
	for _, client := range clients {
		output = append(output, client)
	}
	return output
}
