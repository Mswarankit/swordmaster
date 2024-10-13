package window

import (
	"log"
	"os"
	"runtime"
	"swordmaster/internal/event"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
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
	Dtime   float64
}

func NewWindow(w, h int, title string) *Window {
	outputWindow := Window{
		sx:    1,
		sy:    1,
		Dtime: 0,
	}
	err := glfw.Init()
	if err != nil {
		log.Fatalf("Error initializing GLFW: %v", err)
	}

	glfw.WindowHint(glfw.StencilBits, 8)
	glfw.WindowHint(glfw.DepthBits, 0)

	window, err := glfw.CreateWindow(1280, 720, title, nil, nil)
	if err != nil {
		log.Fatalf("Error creating window: %v", err)
	}
	window.MakeContextCurrent()

	err = gl.Init()
	if err != nil {
		log.Fatalf("Error initializing GL: %v", err)
	}

	glfw.SwapInterval(1)
	gl.Enable(gl.MULTISAMPLE)

	glBackend, err := goglbackend.New(0, 0, 0, 0, nil)
	if err != nil {
		log.Fatalf("Error loading canvas GL assets: %v", err)
	}
	outputWindow.KB = event.NewKeyboard()
	outputWindow.backend = glBackend
	outputWindow.Current = window
	outputWindow.canvas = canvas.New(glBackend)
	outputWindow.init()
	f, _ := outputWindow.canvas.LoadFont(os.Getenv("FONT"))
	outputWindow.canvas.SetFont(f, 18)
	return &outputWindow
}

func (w *Window) Run(fn func(cv *canvas.Canvas, w, h float64)) {
	w.Current.SetCursorPosCallback(event.MousePositionListener)
	w.Current.SetMouseButtonCallback(event.MouseButtonListener)

	pTime := glfw.GetTime()
	targetFrameTime := 1.0 / 60.0

	for !w.Current.ShouldClose() {

		currentTime := glfw.GetTime()
		w.Dtime = currentTime - pTime
		pTime = currentTime

		glfw.PollEvents()
		w.KB.ListenToKeys(w.Current)

		gl.Clear(gl.COLOR_BUFFER_BIT)

		ww, wh := w.Current.GetSize()
		fbw, fbh := w.Current.GetFramebufferSize()

		w.sx = float64(fbw) / float64(ww)
		w.sy = float64(fbh) / float64(wh)

		w.backend.SetBounds(0, 0, fbw, fbh)

		fn(w.canvas, float64(fbw), float64(fbh))

		w.Current.SwapBuffers()

		frameEndTime := glfw.GetTime()
		frameDuration := frameEndTime - pTime

		if frameDuration < targetFrameTime {
			sleepTime := targetFrameTime - frameDuration
			time.Sleep(time.Duration(sleepTime * 1e9))
		}
	}
}

func (w Window) Close() {
	glfw.Terminate()
}

func (w Window) init() {
	runtime.LockOSThread()
}

func (w Window) GetCanvas() *canvas.Canvas {
	return w.canvas
}
