package entity

import (
	"fmt"
	"math"

	"swordmaster/internal/enums"
	"swordmaster/pkg/window"
	"swordmaster/store"

	"github.com/go-gl/glfw/v3.3/glfw"
	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/tfriedel6/canvas"
)

type Bullet struct {
	Origin   string           `json:"origin"`
	Type     enums.BulletType `json:"type"`
	Force    glm.Vec2         `json:"force"`
	Position glm.Vec2         `json:"position"`
	Size     float64          `json:"size"`
}

func NewBullet(name string, bType enums.BulletType, position glm.Vec2, force glm.Vec2) *Bullet {
	bullet := Bullet{
		Force:    force,
		Position: position,
		Type:     bType,
		Size:     10,
		Origin:   fmt.Sprintf("%s_%v", name, glfw.GetTime()),
	}
	store.AddBullet(&bullet)
	return &bullet
}

func (b *Bullet) Setup(w *window.Window) {

}

func (b *Bullet) Draw(cv *canvas.Canvas, w, h float64) {
	cv.SetFillStyle("#F0F")
	cv.BeginPath()
	cv.Arc(b.Position.X(), b.Position.Y(), b.Size, 0, math.Pi*2, false)
	cv.Fill()
}

func (p *Bullet) Update(w *window.Window) {
	p.Position = p.Position.Add(p.Force.Mul(w.Dtime))
}

func (b Bullet) GetOrigin() string {
	return b.Origin
}

func (b Bullet) GetPosition() glm.Vec2 {
	return b.Position
}

func (b Bullet) GetForce() glm.Vec2 {
	return b.Force
}

func (b Bullet) GetSize() float64 {
	return b.Size
}

func (b Bullet) GetType() int {
	return int(b.Type)
}
