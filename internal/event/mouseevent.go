package event

import "github.com/go-gl/glfw/v3.3/glfw"

type MouseEvent struct {
	X         float64
	Y         float64
	listeners map[glfw.MouseButton][]func()
}

var Mouse = MouseEvent{
	listeners: make(map[glfw.MouseButton][]func()),
}

func MousePositionListener(gw *glfw.Window, xpos, ypos float64) {
	Mouse.X, Mouse.Y = xpos, ypos
}

func MouseButtonListener(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	for button, listeners := range Mouse.listeners {
		if glfw.Press == w.GetMouseButton(button) {
			for _, listener := range listeners {
				listener()
			}
		}
	}
}

func AddMouseListener(key glfw.MouseButton, cb func()) {
	listeners := Mouse.listeners[key]
	listeners = append(listeners, cb)
	Mouse.listeners[key] = listeners
}
