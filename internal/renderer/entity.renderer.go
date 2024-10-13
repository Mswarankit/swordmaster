package renderer

import (
	"fmt"
	"swordmaster/pkg/window"
	"swordmaster/types"

	"github.com/go-gl/glfw/v3.3/glfw"
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
	fmt.Printf("DTime: %v, Time: %v\n", r.Window.Dtime, glfw.GetTime())
}

func (r *EntityRenderer) CleanUP() {

}
