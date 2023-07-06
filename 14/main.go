package main

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

type Order struct {
	ID uuid.UUID `json:"id"`
}

func main() {

	uuidEvent := `{
		"id": "da113111-0000-0000-0000-000007876909"
	}`

	var order Order
	if err := json.Unmarshal([]byte(uuidEvent), &order); err != nil {
		log.Fatal(err)
	}

	log.Printf("order: %+v", order)
}
