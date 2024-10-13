package store

import (
	"swordmaster/pkg/window"

	cv "github.com/tfriedel6/canvas"
)

var win *window.Window
var canvas *cv.Canvas

func SetWindow(w *window.Window) {
	win = w
}

func GetWindow() *window.Window {
	return win
}

func SetCanvas(cv *cv.Canvas) {
	canvas = cv
}

func GetCanvas() *cv.Canvas {
	return canvas
}
