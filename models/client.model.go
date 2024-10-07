package models

import "net"

type Client struct {
	Address  *net.UDPAddr
	Name     string
	Position []float64
}

func NewClient(name string, address *net.UDPAddr) *Client {
	return &Client{
		Name:    name,
		Address: address,
	}
}

func (c *Client) SetPosition(newPos []float64) {
	c.Position = newPos
}
