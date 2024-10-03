package window

import (
	"log"
	"runtime"
	"swordmaster/internal/event"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/goglbackend"
)

type Window struct {
	Width   int
	Height  int
	Title   string
	MouseX  float64
	MouseY  float64
	sx      float64
	sy      float64
	canvas  *canvas.Canvas
	backend *goglbackend.GoGLBackend
	Current *glfw.Window
	KB      *event.Keyboard
}

func NewWindow(w, h int, title string) *Window {
	outputWindow := Window{
		sx: 1,
		sy: 1,
	}
	err := glfw.Init()
	if err != nil {
		log.Fatalf("Error initializing GLFW: %v", err)
	}

	glfw.WindowHint(glfw.StencilBits, 8)
	glfw.WindowHint(glfw.DepthBits, 0)

	// create window
	window, err := glfw.CreateWindow(1280, 720, "GLFW Test", nil, nil)
	if err != nil {
		log.Fatalf("Error creating window: %v", err)
	}
	window.MakeContextCurrent()

	// init GL
	err = gl.Init()
	if err != nil {
		log.Fatalf("Error initializing GL: %v", err)
	}

	// set vsync on, enable multisample (if available)
	glfw.SwapInterval(1)
	gl.Enable(gl.MULTISAMPLE)

	// load GL backend
	backend, err := goglbackend.New(0, 0, 0, 0, nil)
	if err != nil {
		log.Fatalf("Error loading canvas GL assets: %v", err)
	}
	outputWindow.KB = event.NewKeyboard()
	outputWindow.backend = backend
	outputWindow.Current = window
	outputWindow.canvas = canvas.New(backend)
	return &outputWindow
}

func (w *Window) Run(fn func(cv *canvas.Canvas, w, h float64)) {
	w.Current.SetCursorPosCallback(func(gw *glfw.Window, xpos, ypos float64) {
		w.MouseX, w.MouseY = xpos*w.sx, ypos*w.sy
	})
	w.Current.SetKeyCallback(w.KB.Listen)
	for !w.Current.ShouldClose() {
		w.Current.MakeContextCurrent()

		// find window size and scaling
		ww, wh := w.Current.GetSize()
		fbw, fbh := w.Current.GetFramebufferSize()
		w.sx = float64(fbw) / float64(ww)
		w.sy = float64(fbh) / float64(wh)

		glfw.PollEvents()

		// set canvas size
		w.backend.SetBounds(0, 0, fbw, fbh)
		// call the run function to do all the drawing
		fn(w.canvas, float64(fbw), float64(fbh))

		// swap back and front buffer
		w.Current.SwapBuffers()
	}
}

func (w Window) Close() {
	glfw.Terminate()
}

func init() {
	runtime.LockOSThread()
}
