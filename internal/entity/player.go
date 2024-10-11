package entity

import (
	"math/rand/v2"
	"swordmaster/models"
	"swordmaster/pkg/utils"
	"swordmaster/pkg/window"
	"swordmaster/store"

	"github.com/go-gl/glfw/v3.3/glfw"
	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/tfriedel6/canvas"
)

type Player struct {
	position glm.Vec2
	size     float64
	name     string
	speed    float64
	color    string
}

func NewPlayer(name string, x, y float64, s float64) *Player {
	newColor := utils.RGBtoHEX(rand.IntN(256), rand.IntN(256), rand.IntN(256), rand.IntN(256))
	return &Player{
		name: name,
		position: glm.Vec2{
			x, y,
		},
		size:  s,
		speed: 5.0,
		color: newColor,
	}
}

func (p *Player) Setup(w *window.Window) {
	w.KB.AddListener(glfw.KeyW, func() {
		p.position = p.position.Add(glm.Vec2{0, -1}.Mul(p.speed))
	})
	w.KB.AddListener(glfw.KeyS, func() {
		p.position = p.position.Add(glm.Vec2{0, 1}.Mul(p.speed))
	})
	w.KB.AddListener(glfw.KeyA, func() {
		p.position = p.position.Add(glm.Vec2{-1, 0}.Mul(p.speed))
	})
	w.KB.AddListener(glfw.KeyD, func() {
		p.position = p.position.Add(glm.Vec2{1, 0}.Mul(p.speed))
	})
}

func (p *Player) Draw(cv *canvas.Canvas, w, h float64) {
	cv.SetFillStyle(p.color)
	if store.GetLink() != nil {
		store.GetLink().Broadcast(&models.Message{
			Kind: "POS",
			Name: p.name,
			Data: []float64{p.position.X(), p.position.Y()},
		})
	}
	cv.FillRect(p.position.X(), p.position.Y(), p.size, p.size)
	cv.FillText(p.name, p.position.X(), p.position.Y()+p.size+16)
	for _, client := range store.GetClients() {
		cv.FillRect(client.PositionX, client.PositionY, p.size, p.size)
	}
}
