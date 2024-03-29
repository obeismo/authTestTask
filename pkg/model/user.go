package model

type User struct {
	ID           string `bson:"_id" json:"id"`
	RefreshToken string `bson:"refresh_token" json:"refresh_token"`
}
