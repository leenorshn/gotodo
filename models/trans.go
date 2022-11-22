package models

import "time"

type Trans struct {
	ID        string    `json:"_id" bson:"_id"`
	From      *Account  `json:"from"`
	Amount    float64   `json:"amount"`
	Operation string    `json:"operation"`
	Date      time.Time `json:"date"`
}
