package entity

import (
	"math"
	"math/rand/v2"
	"strings"
	"swordmaster/internal/enums"
	"swordmaster/internal/event"
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
	Speed    float64
	Dum      float64 `json:"dum"`
}

func NewPlayer(name string, x, y float64, s float64) *Player {
	var mxc = 256
	var mnc = 50
	newColor := utils.RGBtoHEX(mnc+rand.IntN(mxc-mnc), mnc+rand.IntN(mxc-mnc), mnc+rand.IntN(mxc-mnc), mnc+rand.IntN(mxc-mnc))
	return &Player{
		Name:   name,
		Size:   s,
		Speed:  5.0,
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
		p.Position = p.Position.Add(glm.Vec2{0, -1}.Mul(p.Speed))
	})
	w.KB.AddListener(glfw.KeyS, func() {
		p.Position = p.Position.Add(glm.Vec2{0, 1}.Mul(p.Speed))
	})
	w.KB.AddListener(glfw.KeyA, func() {
		p.Position = p.Position.Add(glm.Vec2{-1, 0}.Mul(p.Speed))
	})
	w.KB.AddListener(glfw.KeyD, func() {
		p.Position = p.Position.Add(glm.Vec2{1, 0}.Mul(p.Speed))
	})
	event.AddMouseListener(glfw.MouseButtonLeft, func() {
		mv := glm.Vec2{event.Mouse.X, event.Mouse.Y}
		fv := mv.Sub(p.Position).Normalize().Mul(float64(enums.Normal * 100))
		p.Shout(enums.SHOOT, NewBullet(p.Name, enums.Normal, p.Position, fv))
	})
}

func (p *Player) Shout(kind enums.MessageType, data interface{}) {
	if store.GetLink() != nil {
		store.GetLink().Broadcast(
			kind,
			p.Name,
			io.ToBytes(data),
		)
	}
}

func (p *Player) Draw(cv *canvas.Canvas, w, h float64) {
	cv.SetFillStyle(p.Color)
	p.Shout(enums.POS, p)
	cv.FillRect(p.Position.X(), p.Position.Y(), p.Size, p.Size)
	cv.SetFillStyle("#FFF")
	cv.FillText(p.Name, p.Position.X(), p.Position.Y()+p.Size+18)
	cx := p.Position.X() + p.Size/2
	cy := p.Position.Y() + p.Size/2
	for i := 0.0; i < 2*math.Pi; i += math.Pi / 8 {
		phase := glfw.GetTime()
		x := cx + p.Size*math.Cos(i+phase)
		y := cy + p.Size*math.Sin(i+phase)
		cv.FillRect(x, y, 10, 10)
	}
	cv.SetStrokeStyle("#FFF")
	cv.Stroke()
	dir := glm.Vec2{event.Mouse.X, event.Mouse.Y}.Sub(glm.Vec2{cx, cy})
	dir.Normalize()

	cv.BeginPath()
	cv.MoveTo(cx, cy)
	cv.LineTo(event.Mouse.X, event.Mouse.Y)
	cv.ClosePath()

	var coPlayer Player
	for _, client := range store.GetClients() {
		io.FromBytes(client.Player, &coPlayer)
		cv.SetFillStyle(coPlayer.Color)
		cv.FillRect(coPlayer.Position.X(), coPlayer.Position.Y(), coPlayer.Size, coPlayer.Size)
		cv.SetFillStyle("#FFF")
		cv.FillText(coPlayer.Name, coPlayer.Position.X(), coPlayer.Position.Y()+coPlayer.Size+18)
		cx := coPlayer.Position.X() + coPlayer.Size/2
		cy := coPlayer.Position.Y() + coPlayer.Size/2
		for i := 0.0; i < 2*math.Pi; i += math.Pi / 8 {
			phase := glfw.GetTime()
			x := cx + coPlayer.Size*math.Cos(i+phase)
			y := cy + coPlayer.Size*math.Sin(i+phase)
			cv.FillRect(x, y, 10, 10)
		}
	}
}

func (p *Player) Update(w *window.Window) {
	for origin, bullet := range store.ListBullets() {
		if !strings.HasPrefix(origin, p.Name) && bullet.GetPosition().Sub(p.Position).Len() <= bullet.GetSize()+p.Size {
			p.Shout(enums.HIT, &bullet)
			store.RemoveBullet(bullet.GetOrigin())
		}
	}
}
