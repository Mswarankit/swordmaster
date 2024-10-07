package main

import (
	"swordmaster/internal/entity"
	"swordmaster/internal/gui"
	"swordmaster/internal/renderer"
	wn "swordmaster/pkg/window"
)

var window *wn.Window

func main() {
	window = wn.NewWindow(1366, 768, "MYGL")
	mrenderer := renderer.NewMasterRenderer(window)
	mrenderer.Init([]renderer.Renderer{
		renderer.NewEntityRenderer(
			entity.NewPlayer("Ram", 100, 100, 100),
		),
		renderer.NewUIRenderer(
			window.Current,
			gui.NewNetworkUI(),
		),
	})
	defer mrenderer.CleanUP()
	window.Run(mrenderer.Render)
}
