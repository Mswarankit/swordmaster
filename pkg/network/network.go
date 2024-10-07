package network

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"swordmaster/models"
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
		json.Unmarshal([]byte(buf[:length]), &message)
		if message.Kind == "JOIN" {
			store.AddClient(strings.ToUpper(message.Name), addr)
			fmt.Printf("Position: %v\n", message.Data)
			n.SendMessageTo(&models.Message{
				Kind: "JOIN_SUCCESS",
				Name: "SERVER",
			}, addr)
		}
		if message.Kind == "POS" {
			fmt.Printf("models.Message %v\n", message)
			fmt.Printf("store.ClientIds %v\n", store.ClientIds())
			store.GetClient(strings.ToUpper(message.Name)).SetPosition(message.Data[0], message.Data[1])
		}
	}
}

func (n *UDPNetwork) SendMessageTo(message *models.Message, clientAddr *net.UDPAddr) {
	jd, _ := json.Marshal(message)
	n.conn.WriteToUDP([]byte(jd), clientAddr)
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
		Name: "Manoj",
		Data: []float64{1.0, 2.0},
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
	jsonString := string(buf[:l])
	fmt.Println(jsonString)
	var message models.Message

	// Unmarshal the JSON string into the Message struct
	err = json.Unmarshal([]byte(jsonString), &message)
	if err != nil {
		log.Fatal(err)
	}
	store.AddClient(message.Name, addr)
	go n.listen()
	return output
}

func (n *UDPNetwork) Broadcast(message *models.Message) {
	var wg sync.WaitGroup

	for _, client := range store.GetClients() {
		wg.Add(1)
		go func(client *models.Client) {
			defer wg.Done()
			n.SendMessageTo(message, client.Address)
		}(client)
	}

	wg.Wait() // Wait for all goroutines to finish
}

func (n *UDPNetwork) Close() {
	n.conn.Close()
}
