package models

import (
	"fmt"
	"net"
)

type Client struct {
	Address   *net.UDPAddr
	Name      string
	PositionX float64
	PositionY float64
}

func NewClient(name string, address *net.UDPAddr) *Client {
	return &Client{
		Name:      name,
		Address:   address,
		PositionX: -1000,
		PositionY: -1000,
	}
}

func (c *Client) SetPosition(x, y float64) {
	fmt.Println(c, x, y)
	c.PositionX = x
	c.PositionY = y
}
