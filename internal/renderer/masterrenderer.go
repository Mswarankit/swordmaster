package renderer

import (
	"swordmaster/pkg/window"

	"github.com/tfriedel6/canvas"
)

type Renderer interface {
	Setup(*window.Window)
	Render(*canvas.Canvas, float64, float64)
}

type MasterRenderer struct {
	w         *window.Window
	renderers []Renderer
}

func NewMasterRenderer(w *window.Window) *MasterRenderer {
	return &MasterRenderer{
		w: w,
	}
}

func (mr *MasterRenderer) Init(renderers []Renderer) {
	mr.renderers = renderers
	for _, renderers := range mr.renderers {
		renderers.Setup(mr.w)
	}
}

func (mr *MasterRenderer) Render(cv *canvas.Canvas, w, h float64) {
	cv.SetFillStyle("#000")
	cv.FillRect(0, 0, w, h)
	for _, renderers := range mr.renderers {
		renderers.Render(cv, w, h)
	}
}
