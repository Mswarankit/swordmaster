package main

import (
	wn "swordmaster/pkg/window"

	"github.com/tfriedel6/canvas"
)

var window *wn.Window

func main() {

	window = wn.NewWindow(1366, 768, "MYGL")
	window.Run(run)
}

func run(cv *canvas.Canvas, w, h float64) {
	cv.SetFillStyle("#000")
	cv.FillRect(0, 0, w, h)
	cv.SetFillStyle("#00F")
	cv.FillRect(w*0.25, h*0.25, w*0.5, h*0.5)
	cv.SetStrokeStyle("#0F0")
	cv.StrokeRect(window.MouseX-32, window.MouseY-32, 64, 64)
}
