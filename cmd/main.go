package main

import (
	"fmt"
	"log"
	"os"
	"swordmaster/internal/entity"
	"swordmaster/internal/gui"
	"swordmaster/internal/renderer"
	wn "swordmaster/pkg/window"
	"swordmaster/store"
	"swordmaster/types"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	window := wn.NewWindow(1366, 768, fmt.Sprintf("MYGL - %v", os.Getenv("MY_NAME")))
	store.SetWindow(window)
	store.SetCanvas(window.GetCanvas())
	mrenderer := renderer.NewMasterRenderer(window)
	mrenderer.Init([]types.Renderer{
		renderer.NewEntityRenderer(
			window,
			entity.NewPlayer(os.Getenv("MY_NAME"), 100, 100, 30),
		),
		renderer.NewBulletRenderer(window),
		renderer.NewUIRenderer(
			window,
			gui.NewNetworkUI(),
			gui.NewStatsUI(),
		),
	})
	defer mrenderer.CleanUP()
	window.Run(mrenderer.Render)
}
