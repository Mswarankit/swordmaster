package renderer

import (
	"swordmaster/pkg/window"
	"swordmaster/store"

	"github.com/tfriedel6/canvas"
)

type BulletRenderer struct {
	Window *window.Window
}

func NewBulletRenderer(w *window.Window) *BulletRenderer {
	return &BulletRenderer{
		Window: w,
	}
}

func (r *BulletRenderer) Setup() {
	for _, e := range store.ListBullets() {
		e.Setup(r.Window)
	}
}

func (r *BulletRenderer) Render(cv *canvas.Canvas, w, h float64) {
	for _, e := range store.ListBullets() {
		e.Update(r.Window)
		e.Draw(cv, w, h)
		if e.GetPosition().X() > w || e.GetPosition().X() < 0 || e.GetPosition().Y() > h || e.GetPosition().Y() < 0 {
			store.RemoveBullet(e.GetOrigin())
		}
	}
}

func (r *BulletRenderer) CleanUP() {

}
