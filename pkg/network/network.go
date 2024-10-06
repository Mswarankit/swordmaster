package network

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"swordmaster/configs"
	"sync"
)

const DEFAULT_PORT = 9211

type Network struct {
	conn      *net.UDPConn
	myAddress net.Addr
}

type Message struct {
	Kind string    `json:"kind"`
	Name string    `json:"name"`
	Data []float64 `json:"data"`
}

func NewNetwork() *Network {
	return &Network{}
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

func (n Network) GetAddress() string {
	addrs, err := net.InterfaceAddrs()
	defAddr := fmt.Sprintf("http://localhost:%v", DEFAULT_PORT)
	if err != nil {
		return defAddr
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return fmt.Sprintf("http://%v:%v", ipNet.IP.String(), DEFAULT_PORT)
			}
		}
	}

	return defAddr
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
			configs.AddClient(message.Name, addr)
			fmt.Printf("Position: %v\n", message.Data)
			n.SendMessageTo(&Message{
				Kind: "JOIN_SUCCESS",
				Name: "SERVER",
			}, addr)
		}
		if message.Kind == "POS" {
			fmt.Printf("Message %v\n", message)
		}
	}
}

func (n *Network) SendMessageTo(message *Message, clientAddr *net.UDPAddr) {
	jd, _ := json.Marshal(message)
	n.conn.WriteToUDP([]byte(jd), clientAddr)
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

func (n *Network) Broadcast(message *Message) {
	var wg sync.WaitGroup

	for _, client := range configs.ClientAddresses() {
		wg.Add(1)
		go func(client *net.UDPAddr) {
			defer wg.Done()
			n.SendMessageTo(message, client)
		}(client)
	}

	wg.Wait() // Wait for all goroutines to finish
}

func (n *Network) Close() {
	n.conn.Close()
}
