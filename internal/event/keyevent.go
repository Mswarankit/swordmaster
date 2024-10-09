package event

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Keyboard struct {
	listeners map[glfw.Key][]func()
}

func NewKeyboard() *Keyboard {
	return &Keyboard{
		listeners: make(map[glfw.Key][]func()),
	}
}

func (keyboard *Keyboard) ListenToKeys(w *glfw.Window) {
	for key, listeners := range keyboard.listeners {
		if glfw.Press == w.GetKey(key) {
			for _, listener := range listeners {
				listener()
			}
		}
	}
}

func (keyboard *Keyboard) AddListener(key glfw.Key, cb func()) {
	listeners := keyboard.listeners[key]
	listeners = append(listeners, cb)
	keyboard.listeners[key] = listeners
}
