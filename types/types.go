package types

import (
	"net"
	"swordmaster/internal/enums"
	"swordmaster/internal/models"
	"swordmaster/pkg/window"

	"github.com/tfriedel6/canvas"

	glm "github.com/go-gl/mathgl/mgl64"
)

type Network interface {
	CreateServer(addrs ...string)
	JoinServer(addrs string) bool
	SendMessageTo(message *models.Message, clientAddr *net.UDPAddr)
	GetAddress() string
	Broadcast(kind enums.MessageType, name string, message []byte)
	IsServer() bool
	Close()
}

type Entity interface {
	Setup(*window.Window)
	Draw(*canvas.Canvas, float64, float64)
	Update(*window.Window)
}

type Bullet interface {
	Entity
	GetOrigin() string
	GetPosition() glm.Vec2
	GetForce() glm.Vec2
	GetSize() float64
}

type Renderer interface {
	Setup()
	Render(*canvas.Canvas, float64, float64)
	CleanUP()
}
