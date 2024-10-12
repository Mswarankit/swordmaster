package io

import (
	"encoding/json"
	"log"
)

func ToBytes(data interface{}) []byte {
	byteArray, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON to byte array: %v", err)
	}
	return byteArray
}

func FromBytes(data []byte, n interface{}) {
	if len(data) > 0 {
		err := json.Unmarshal(data, n)
		if err != nil {
			log.Fatalf("Error marshalling Byte to JSON : %v", err)
		}
	}
}
