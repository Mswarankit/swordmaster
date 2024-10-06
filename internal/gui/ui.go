package gui

import "github.com/tfriedel6/canvas"

type UI interface {
	Setup()
	Draw(*canvas.Canvas, float64, float64)
}
