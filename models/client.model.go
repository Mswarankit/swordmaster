package models

import "net"

type Client struct {
	Address   *net.UDPAddr
	Name      string
	PositionX float64
	PositionY float64
}

func NewClient(name string, address *net.UDPAddr) *Client {
	return &Client{
		Name:    name,
		Address: address,
	}
}

func (c *Client) SetPosition(newPos []float64) {
	c.PositionX = newPos[0]
	c.PositionY = newPos[0]
}
