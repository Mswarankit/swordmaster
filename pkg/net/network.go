package net

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

const DEFAULT_PORT = 9211

type Network struct {
	conn      *net.UDPConn
	myAddress net.Addr
	myClients map[string]*net.UDPAddr
}

type Message struct {
	Kind string    `json:"kind"`
	Name string    `json:"name"`
	Data []float64 `json:"data"`
}

func NewNetwork() *Network {
	return &Network{
		myClients: make(map[string]*net.UDPAddr),
	}
}

func (n *Network) CreateServer(adrs ...string) {
	var adr string
	if len(adrs) > 0 {
		adr = adrs[0]
	} else {
		adr = fmt.Sprintf("0.0.0.0: %d", DEFAULT_PORT)
	}
	addr, err := net.ResolveUDPAddr("udp", adr)
	if err != nil {
		log.Fatal(err)
	}
	n.myAddress = addr
	ln, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	n.conn = ln
	go n.listen()
}

func (n *Network) listen() {
	buf := make([]byte, 4096)
	for {
		length, addr, err := n.conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		var message Message
		json.Unmarshal([]byte(buf[:length]), &message)
		if message.Kind == "JOIN" {
			fmt.Println(n.myClients, addr.String())
			n.myClients[addr.String()] = addr
			fmt.Printf("Position: %v\n", message.Data)
			jd, _ := json.Marshal(Message{
				Kind: "JOIN_SUCCESS",
				Name: "SERVER",
			})
			n.conn.WriteToUDP([]byte(jd), addr)
		}
		if message.Kind == "POS" {
			fmt.Printf("Message %v\n", message)
		}
	}
}

func (n *Network) Join(serverAddress string) {
	addr, err := net.ResolveUDPAddr("udp", serverAddress)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	n.conn = conn
}

func (n *Network) Send(message *Message) {
	buf := make([]byte, 1024)
	for {
		jsonData, err := json.Marshal(Message{
			Kind: "POS",
			Name: "Client",
			Data: []float64{1.0, 2.0, 3.0},
		})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		_, err = n.conn.Write([]byte(jsonData))
		if err != nil {
			log.Fatal(err)
		}
		n, err := n.conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf[:n]))
		time.Sleep(time.Second)
	}
}

func (n *Network) Close() {
	n.conn.Close()
}
