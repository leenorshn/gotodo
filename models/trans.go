package models

type Trans struct {
	ID        string   `json:"_id" bson:"_id"`
	From      *Account `json:"from"`
	To        *Account `json:"to"`
	Amount    float64  `json:"amount"`
	Operation string   `json:"operation"`
	Date      float64  `json:"date"`
}
