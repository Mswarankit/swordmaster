package gui

import (
	"fmt"
	"swordmaster/configs"
	"swordmaster/pkg/network"

	"github.com/micahke/imgui-go"
	"github.com/tfriedel6/canvas"
)

type NetworkUI struct {
	server *network.Network
}

func NewNetworkUI() *NetworkUI {
	return &NetworkUI{}
}

func (ui *NetworkUI) Setup() {

}

func (ui *NetworkUI) Draw(cv *canvas.Canvas, w, h float64) {
	imgui.Begin("Network Info")
	clicked := imgui.Button("Create Server")
	if clicked && ui.server == nil {
		ui.server = network.NewNetwork()
		ui.server.CreateServer()
	}
	if ui.server != nil {
		imgui.Text(ui.server.GetAddress())
	}
	imgui.LabelText("Active Clients", fmt.Sprintf("%d", configs.ClientCount()))
	if configs.ClientCount() > 0 {
		for _, c := range configs.ClientIds() {
			imgui.Text(c)
		}
	}
	imgui.End()
}
