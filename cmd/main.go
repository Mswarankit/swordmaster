package main

import (
	"fmt"
	"log"
	"os"
	"swordmaster/internal/entity"
	"swordmaster/internal/gui"
	"swordmaster/internal/renderer"
	wn "swordmaster/pkg/window"

	"github.com/joho/godotenv"
)

var window *wn.Window

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	window = wn.NewWindow(1366, 768, fmt.Sprintf("MYGL - %v", os.Getenv("MY_NAME")))
	mrenderer := renderer.NewMasterRenderer(window)
	mrenderer.Init([]renderer.Renderer{
		renderer.NewEntityRenderer(
			entity.NewPlayer(os.Getenv("MY_NAME"), 100, 100, 100),
		),
		renderer.NewUIRenderer(
			window.Current,
			gui.NewNetworkUI(),
		),
	})
	defer mrenderer.CleanUP()
	window.Run(mrenderer.Render)
}
