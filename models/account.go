package models

type Account struct {
	ID            string  `json:"_id" bson:"_id"`
	Balance       float64 `json:"balance"`
	Dette         float64 `json:"dette"`
	Gain          float64 `json:"gain"`
	TotalWithDraw float64 `json:"totalWithDraw"`
	Round         *int    `json:"round"`
	CreatedAt     float64 `json:"createdAt"`
	UpdatedAt     float64 `json:"updatedAt"`
	User          *User   `json:"user"`
}
