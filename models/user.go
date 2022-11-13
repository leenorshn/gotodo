package models

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
