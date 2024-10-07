package models

type Message struct {
	Kind string    `json:"kind"`
	Name string    `json:"name"`
	Data []float64 `json:"data"`
}
