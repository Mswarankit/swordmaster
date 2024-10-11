package models

type PlayerState struct {
	Position []float64 `json:"position"`
}

type Message struct {
	Kind string      `json:"kind"`
	Name string      `json:"name"`
	Data PlayerState `json:"data"`
}
