package models

type Motar struct {
	ID          string `json:"_id" bson:"_id"`
	Association string `json:"association"`
	Numero      string `json:"numero"`
	Parking     string `json:"parking"`
	City        string `json:"city"`
	Owner       *User  `json:"owner"`
}
