package main

import (
	"swordmaster/internal/entity"
	"swordmaster/internal/renderer"
	"swordmaster/pkg/net"
	wn "swordmaster/pkg/window"
)

var window *wn.Window

func main() {
	window = wn.NewWindow(1366, 768, "MYGL")
	mrenderer := renderer.NewMasterRenderer(window)
	mrenderer.Init([]renderer.Renderer{
		renderer.NewEntityRenderer(
			entity.NewPlayer(100, 100, 100),
			entity.NewPlayer(100, 400, 100),
		),
	})
	network := net.NewNetwork()
	network.CreateServer()
	defer network.Close()
	window.Run(mrenderer.Render)
}
