package network

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"swordmaster/internal/models"
	"swordmaster/pkg/io"
	"swordmaster/store"
	"swordmaster/types"
	"sync"
)

const DEFAULT_PORT = 9211

type UDPNetwork struct {
	conn      *net.UDPConn
	myAddress net.Addr
}

func NewNetwork() types.Network {
	return &UDPNetwork{}
}

func (n *UDPNetwork) CreateServer(adrs ...string) {
	var adr string
	if len(adrs) > 0 {
		adr = adrs[0]
	} else {
		adr = fmt.Sprintf("0.0.0.0:%d", DEFAULT_PORT)
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

func (n UDPNetwork) GetAddress() string {
	defAddr := fmt.Sprintf("http://localhost:%v", DEFAULT_PORT)

	// Get a list of all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		return defAddr
	}

	// Iterate over the interfaces to find the one associated with the Wi-Fi hotspot
	for _, iface := range interfaces {
		// Check if the interface is up and not loopback
		if iface.Flags&net.FlagUp != 0 && !(iface.Flags&net.FlagLoopback != 0) {
			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}

			// Check each address associated with this interface
			for _, addr := range addrs {
				if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.To4() != nil {
					// Check if the interface name contains "Wi-Fi" or "Wireless"
					if strings.Contains(strings.ToLower(iface.Name), "wifi") || strings.Contains(strings.ToLower(iface.Name), "wireless") {
						return fmt.Sprintf("http://%v:%v", ipNet.IP.String(), DEFAULT_PORT)
					}
				}
			}
		}
	}

	// Return default address if no suitable address is found
	return defAddr
}

func (n *UDPNetwork) listen() {
	buf := make([]byte, 4096)
	for {
		length, addr, err := n.conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		var message models.Message
		io.FromBytes(buf[:length], &message)
		if message.Kind == "JOIN" {
			store.AddClient(strings.ToUpper(message.Name), nil, addr)
			fmt.Printf("Position: %v\n", message.Data)
			n.SendMessageTo(&models.Message{
				Kind: "JOIN_SUCCESS",
				Name: os.Getenv("MY_NAME"),
			}, addr)
		}
		if message.Kind == "POS" {
			store.GetClient(strings.ToUpper(message.Name)).SetPlayer(message.Data)
		}
	}
}

func (n *UDPNetwork) SendMessageTo(message *models.Message, clientAddr *net.UDPAddr) {
	out, err := n.conn.Write(io.ToBytes(message))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
	// n.conn.WriteToUDP([]byte(jd), clientAddr)
}

func (n *UDPNetwork) JoinServer(serverAddress string) bool {
	addr, err := net.ResolveUDPAddr("udp", serverAddress)
	output := true
	if err != nil {
		log.Fatal(err)
		output = false
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
		output = false
	}
	n.conn = conn
	jsonData, err := json.Marshal(models.Message{
		Kind: "JOIN",
		Name: os.Getenv("MY_NAME"),
		Data: nil,
	})
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Write([]byte(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	l, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	var message models.Message
	io.FromBytes(buf[:l], &message)

	store.AddClient(strings.ToUpper(message.Name), message.Data, addr)
	go n.listen()
	return output
}

func (n *UDPNetwork) Broadcast(kind string, name string, data []byte) {
	var wg sync.WaitGroup
	for _, client := range store.GetClients() {
		wg.Add(1)
		go func(client *models.Client) {
			defer wg.Done()
			n.SendMessageTo(&models.Message{
				Kind: kind,
				Name: name,
				Data: data,
			}, client.Address)
		}(client)
	}

	wg.Wait() // Wait for all goroutines to finish
}

func (n *UDPNetwork) Close() {
	n.conn.Close()
}
