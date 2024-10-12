package models

import (
	"net"
)

type Client struct {
	Address *net.UDPAddr
	Name    string
	Player  []byte
}

func NewClient(name string, address *net.UDPAddr, player []byte) *Client {
	return &Client{
		Name:    name,
		Address: address,
		Player:  player,
	}
}

func (c *Client) SetPlayer(player []byte) {
	c.Player = player
}
