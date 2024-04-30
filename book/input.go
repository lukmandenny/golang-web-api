package book

import "encoding/json"

type BookInput struct {
	Title    string      `json:"title" binding:"required"`
	Harga    json.Number `json:"harga" binding:"required,number"`
	SubTitle string      `json:"sub_title"` //directive
}
