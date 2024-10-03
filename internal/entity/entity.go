package entity

import (
	"swordmaster/pkg/window"

	"github.com/tfriedel6/canvas"
)

type Entity interface {
	Setup(*window.Window)
	Draw(*canvas.Canvas, float64, float64)
}
