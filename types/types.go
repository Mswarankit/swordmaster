package types

import (
	"net"
	"swordmaster/models"
)

type Network interface {
	CreateServer(addrs ...string)
	JoinServer(addrs string) bool
	SendMessageTo(message *models.Message, clientAddr *net.UDPAddr)
	GetAddress() string
	Broadcast(message *models.Message)
	Close()
}
