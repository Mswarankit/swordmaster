package types

import (
	"net"
	"swordmaster/internal/models"
)

type Network interface {
	CreateServer(addrs ...string)
	JoinServer(addrs string) bool
	SendMessageTo(message *models.Message, clientAddr *net.UDPAddr)
	GetAddress() string
	Broadcast(kind string, name string, message []byte)
	Close()
}
