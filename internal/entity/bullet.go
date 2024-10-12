package entity

import (
	"swordmaster/types"

	glm "github.com/go-gl/mathgl/mgl64"
)

type Bullet struct {
	Type     types.BulletType
	Force    glm.Vec2
	Position glm.Vec2
}

func NewBullet(name string, damage float64, force glm.Vec2, position glm.Vec2) *Bullet {
	return &Bullet{
		Force:    force,
		Position: position,
	}
}
