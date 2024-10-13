package models

import "swordmaster/internal/enums"

type Message struct {
	Kind enums.MessageType `json:"kind"`
	Name string            `json:"name"`
	Data []byte            `json:"data"`
}
