package renderer

import (
	"swordmaster/internal/gui"
	"swordmaster/pkg/window"

	"github.com/go-gl/glfw/v3.3/glfw"
	backend "github.com/micahke/glfw_imgui_backend"
	"github.com/micahke/imgui-go"
	"github.com/tfriedel6/canvas"
)

type UIRenderer struct {
	uis     []gui.UI
	Impl    *backend.ImguiGlfw3
	Context *imgui.Context
	Window  *window.Window
}

func NewUIRenderer(window *window.Window, e ...gui.UI) *UIRenderer {
	ui := UIRenderer{
		uis:     e,
		Context: imgui.CreateContext(nil),
		Window:  window,
	}
	io := imgui.CurrentIO()
	ui.Window.Current.SetCharCallback(func(w *glfw.Window, char rune) {
		io.AddInputCharacters(string(char))
	})
	ui.Window.Current.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Press {
			io.KeyPress(int(key))
		} else if action == glfw.Release {
			io.KeyRelease(int(key))
		}
	})
	ui.Impl = backend.ImguiGlfw3Init(ui.Window.Current, io)
	return &ui
}

func (r *UIRenderer) Setup() {
	for _, e := range r.uis {
		e.Setup()
	}
}

func (r *UIRenderer) Render(cv *canvas.Canvas, w, h float64) {
	r.Impl.NewFrame()
	for _, ui := range r.uis {
		ui.Draw(cv, w, h)
	}
	imgui.Render()
	r.Impl.Render(imgui.RenderedDrawData())
}

func (r *UIRenderer) CleanUP() {
	defer r.Context.Destroy()
	defer r.Impl.Shutdown()
}
