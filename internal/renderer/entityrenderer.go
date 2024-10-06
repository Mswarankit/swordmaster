package renderer

import (
	"swordmaster/internal/entity"
	"swordmaster/pkg/window"

	"github.com/tfriedel6/canvas"
)

type EntityRenderer struct {
	entities []entity.Entity
}

func NewEntityRenderer(e ...entity.Entity) *EntityRenderer {
	return &EntityRenderer{
		entities: e,
	}
}

func (r *EntityRenderer) Setup(w *window.Window) {
	for _, e := range r.entities {
		e.Setup(w)
	}
}

func (r *EntityRenderer) Render(cv *canvas.Canvas, w, h float64) {
	for _, e := range r.entities {
		e.Draw(cv, w, h)
	}
}

func (r *EntityRenderer) CleanUP() {

}
