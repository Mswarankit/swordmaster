package main

import (
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	backend "github.com/micahke/glfw_imgui_backend"
	"github.com/micahke/imgui-go"
)

var showDemoWindow bool

func init() {
	runtime.LockOSThread()
}

type GUI struct {
	Impl    *backend.ImguiGlfw3
	Context *imgui.Context
}

func NewGui(window *glfw.Window) *GUI {
	showDemoWindow = true
	return &GUI{
		Context: imgui.CreateContext(nil),
		Impl:    backend.ImguiGlfw3Init(window, imgui.CurrentIO()),
	}
}

func (g *GUI) Draw() {
	g.Impl.NewFrame()
	imgui.Text("Hello, world!")
	imgui.ShowDemoWindow(&showDemoWindow)
	imgui.Render()
	g.Impl.Render(imgui.RenderedDrawData())
}

func (g *GUI) Close() {
	defer g.Context.Destroy()
	defer g.Impl.Shutdown()
}

var gui *GUI

func main() {

	// Initialize GLFW through go-gl/glfw
	if err := glfw.Init(); err != nil {
		panic("Error initializing GLFW")
	}
	defer glfw.Terminate()

	// GLFW setup
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// Initialize window through go-gl/glfw
	window, win_err := glfw.CreateWindow(920, 540, "Hello, world!", nil, nil)
	if win_err != nil {
		panic("Error creating window")
	}
	gui = NewGui(window)
	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	if err := gl.Init(); err != nil {
		panic("Error initializing OpenGL")
	}

	window.SetKeyCallback(KeyCallback)

	for !window.ShouldClose() {
		glfw.PollEvents()
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gui.Draw()
		window.SwapBuffers()
	}
	defer gui.Close()
}

func KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

}
