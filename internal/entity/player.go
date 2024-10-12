package entity

import (
	"math/rand/v2"
	"swordmaster/pkg/io"
	"swordmaster/pkg/utils"
	"swordmaster/pkg/window"
	"swordmaster/store"

	"github.com/go-gl/glfw/v3.3/glfw"
	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/tfriedel6/canvas"
)

type Player struct {
	Position glm.Vec2 `json:"position"`
	Size     float64  `json:"size"`
	Name     string   `json:"name"`
	Color    string   `json:"color"`
	Health   int      `json:"health"`
	speed    float64
}

func NewPlayer(name string, x, y float64, s float64) *Player {
	newColor := utils.RGBtoHEX(rand.IntN(256), rand.IntN(256), rand.IntN(256), rand.IntN(256))
	return &Player{
		Name:   name,
		Size:   s,
		speed:  5.0,
		Health: 100,
		Color:  newColor,
		Position: glm.Vec2{
			x, y,
		},
	}
}

func (p *Player) SetPosition(pos glm.Vec2) {
	p.Position = pos
}

func (p *Player) SetColor(color string) {
	p.Color = color
}

func (p *Player) Setup(w *window.Window) {
	w.KB.AddListener(glfw.KeyW, func() {
		p.Position = p.Position.Add(glm.Vec2{0, -1}.Mul(p.speed))
	})
	w.KB.AddListener(glfw.KeyS, func() {
		p.Position = p.Position.Add(glm.Vec2{0, 1}.Mul(p.speed))
	})
	w.KB.AddListener(glfw.KeyA, func() {
		p.Position = p.Position.Add(glm.Vec2{-1, 0}.Mul(p.speed))
	})
	w.KB.AddListener(glfw.KeyD, func() {
		p.Position = p.Position.Add(glm.Vec2{1, 0}.Mul(p.speed))
	})
}

func (p *Player) Draw(cv *canvas.Canvas, w, h float64) {
	cv.SetFillStyle(p.Color)
	if store.GetLink() != nil {
		store.GetLink().Broadcast(
			"POS",
			p.Name,
			io.ToBytes(p),
		)
	}
	cv.FillRect(p.Position.X(), p.Position.Y(), p.Size, p.Size)
	cv.FillText(p.Name, p.Position.X(), p.Position.Y()+p.Size+18)
	var coPlayer Player
	for _, client := range store.GetClients() {
		io.FromBytes(client.Player, &coPlayer)
		cv.FillRect(coPlayer.Position.X(), coPlayer.Position.Y(), coPlayer.Size, coPlayer.Size)
		cv.FillText(coPlayer.Name, coPlayer.Position.X(), coPlayer.Position.Y()+coPlayer.Size+18)
	}
}
