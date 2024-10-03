package entity

import (
	"swordmaster/pkg/window"

	"github.com/go-gl/glfw/v3.3/glfw"
	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/tfriedel6/canvas"
)

type Player struct {
	position glm.Vec2
	size     float64
}

func NewPlayer(x, y float64, s float64) *Player {
	return &Player{
		position: glm.Vec2{
			x, y,
		},
		size: s,
	}
}

func (p *Player) Setup(w *window.Window) {
	w.KB.AddListener(glfw.KeyW, func() {
		p.position = p.position.Add(glm.Vec2{0, -1})
	})
	w.KB.AddListener(glfw.KeyS, func() {
		p.position = p.position.Add(glm.Vec2{0, 1})
	})
	w.KB.AddListener(glfw.KeyA, func() {
		p.position = p.position.Add(glm.Vec2{-1, 0})
	})
	w.KB.AddListener(glfw.KeyD, func() {
		p.position = p.position.Add(glm.Vec2{1, 0})
	})
}

func (p *Player) Draw(cv *canvas.Canvas, w, h float64) {
	cv.SetFillStyle("#00F")
	cv.FillRect(p.position.X(), p.position.Y(), p.size, p.size)
}
