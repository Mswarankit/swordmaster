package shapes

import (
	"math"
	"swordmaster/store"

	glm "github.com/go-gl/mathgl/mgl64"
)

func Circle(pos glm.Vec2, radius float64, color string) {
	cv := store.GetCanvas()
	cv.Save()
	cv.SetFillStyle(color)
	cv.SetStrokeStyle(color)
	cv.BeginPath()
	cv.Arc(pos.X(), pos.Y(), radius, 0, math.Pi*2, false)
	cv.Fill()
	cv.Restore()
}

func Rect(pos glm.Vec2, w, h float64, color string) {
	x := pos.X() - w/2
	y := pos.Y() - h/2
	cv := store.GetCanvas()
	cv.Save()
	cv.SetFillStyle(color)
	cv.SetStrokeStyle(color)
	cv.BeginPath()
	cv.FillRect(x, y, w, h)
	cv.Fill()
	cv.Restore()
}
