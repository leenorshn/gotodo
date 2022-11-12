// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

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

type Auth struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type Motar struct {
	ID          string `json:"_id" bson:"_id"`
	Association string `json:"association"`
	Numero      string `json:"numero"`
	Parking     string `json:"parking"`
	City        string `json:"city"`
	Owner       *User  `json:"owner"`
}

type NewMotor struct {
	Text        string `json:"text"`
	User        string `json:"user"`
	Association string `json:"association"`
	Numero      string `json:"numero"`
	Parking     string `json:"parking"`
	City        string `json:"city"`
}

type NewTransaction struct {
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    float64 `json:"amount"`
	Operation string  `json:"operation"`
}

type NewUser struct {
	Name          string  `json:"name"`
	Phone         string  `json:"phone"`
	Password      string  `json:"password"`
	DateNaissance int     `json:"dateNaissance"`
	Avatar        *string `json:"avatar"`
	Address       string  `json:"address"`
}

type Trans struct {
	ID        string   `json:"_id" bson:"_id"`
	From      *Account `json:"from"`
	To        *Account `json:"to"`
	Amount    float64  `json:"amount"`
	Operation string   `json:"operation"`
	Date      float64  `json:"date"`
}

type UpdateMotor struct {
	Text        *string `json:"text"`
	Association *string `json:"association"`
	Numero      *string `json:"numero"`
	Parking     *string `json:"parking"`
	City        *string `json:"city"`
}

type UpdateUser struct {
	Name          *string `json:"name"`
	DateNaissance *int    `json:"dateNaissance"`
	Avatar        *string `json:"avatar"`
	Address       *string `json:"address"`
}

type User struct {
	ID            string   `json:"_id" bson:"_id"`
	Name          string   `json:"name"`
	Phone         string   `json:"phone"`
	Password      string   `json:"password"`
	DateNaissance int      `json:"dateNaissance"`
	Avatar        string   `json:"avatar"`
	Address       string   `json:"address"`
	Account       *Account `json:"account"`
}
