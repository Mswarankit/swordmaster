package renderer

import (
	"swordmaster/pkg/window"
	"swordmaster/types"

	"github.com/tfriedel6/canvas"
)

type EntityRenderer struct {
	entities []types.Entity
	Window   *window.Window
}

func NewEntityRenderer(w *window.Window, e ...types.Entity) *EntityRenderer {
	return &EntityRenderer{
		entities: e,
		Window:   w,
	}
}

func (r *EntityRenderer) Setup() {
	for _, e := range r.entities {
		e.Setup(r.Window)
	}
}

func (r *EntityRenderer) Render(cv *canvas.Canvas, w, h float64) {
	for _, e := range r.entities {
		e.Update(r.Window)
		e.Draw(cv, w, h)
	}
}

func (r *EntityRenderer) CleanUP() {

}
