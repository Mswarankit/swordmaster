package gui

import (
	"fmt"
	"swordmaster/store"

	"github.com/micahke/imgui-go"
	"github.com/tfriedel6/canvas"
)

type StatsUI struct {
}

func NewStatsUI() *StatsUI {
	return &StatsUI{}
}

func (ui *StatsUI) Setup() {

}

func (ui *StatsUI) Draw(cv *canvas.Canvas, w, h float64) {
	{
		imgui.Begin("Bullet Info")
		imgui.LabelText("Active Bullet Count", fmt.Sprintf("%d", store.BulletCount()))
		if store.BulletCount() > 0 {
			for _, b := range store.ListBullets() {
				imgui.Text(b.GetOrigin())
			}
		}
		imgui.End()
	}
}
