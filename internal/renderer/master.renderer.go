package renderer

import (
	"swordmaster/pkg/window"
	"swordmaster/types"

	"github.com/tfriedel6/canvas"
)

type MasterRenderer struct {
	w         *window.Window
	renderers []types.Renderer
}

func NewMasterRenderer(w *window.Window) *MasterRenderer {
	return &MasterRenderer{
		w: w,
	}
}

func (mr *MasterRenderer) Init(renderers []types.Renderer) {
	mr.renderers = renderers
	for _, renderers := range mr.renderers {
		renderers.Setup()
	}
}

func (mr *MasterRenderer) Render(cv *canvas.Canvas, w, h float64) {
	cv.SetFillStyle("#000")
	cv.FillRect(0, 0, w, h)
	for _, renderers := range mr.renderers {
		renderers.Render(cv, w, h)
	}
}

func (mr *MasterRenderer) CleanUP() {
	for _, renderers := range mr.renderers {
		renderers.CleanUP()
	}
}
