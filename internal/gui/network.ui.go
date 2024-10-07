package gui

import (
	"fmt"
	"os"
	"swordmaster/pkg/network"
	"swordmaster/store"

	"github.com/micahke/imgui-go"
	"github.com/tfriedel6/canvas"
)

var serverAddress string = os.Getenv("SERVER_ADDRESS")

type NetworkUI struct {
}

func NewNetworkUI() *NetworkUI {
	return &NetworkUI{}
}

func (ui *NetworkUI) Setup() {

}

func (ui *NetworkUI) Draw(cv *canvas.Canvas, w, h float64) {
	{
		imgui.Begin("Network Info")
		clicked := imgui.Button("Create Server")
		if clicked && store.GetLink() == nil {
			store.SetLink(network.NewNetwork())
			store.GetLink().CreateServer()
		}
		if store.GetLink() != nil {
			imgui.Text(store.GetLink().GetAddress())
		}
		imgui.LabelText("Active Clients", fmt.Sprintf("%d", store.ClientCount()))
		if store.ClientCount() > 0 {
			for _, c := range store.ClientIds() {
				imgui.Text(c)
			}
		}
		imgui.End()
	}

	{
		imgui.Begin("Server")
		// imgui.InputText("Server Addresss", &serverAddress)
		clicked := imgui.Button("Join Server")
		joined := false
		if clicked && len(serverAddress) > 4 {
			fmt.Printf("Server Address: %v\n", serverAddress)
			store.SetLink(network.NewNetwork())
			joined = store.GetLink().JoinServer(serverAddress)
		}

		if joined {
			imgui.Text(fmt.Sprintf("Joined server %v", serverAddress))
		}
		imgui.End()
	}
}
