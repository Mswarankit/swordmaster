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

func (keyboard *Keyboard) Listen(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	for _, listener := range keyboard.listeners[key] {
		listener()
	}
}

func (keyboard *Keyboard) AddListener(key glfw.Key, cb func()) {
	listeners := keyboard.listeners[key]
	listeners = append(listeners, cb)
	keyboard.listeners[key] = listeners
}
