package book

import "encoding/json"

type BookRequestType struct {
	Title       string      `json:"title" binding:"required"`
	SubTitle    string      `json:"sub_title" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Discount    int         `json:"discount" binding:"number"`
	Rating      int         `json:"rating" binding:"number"`
}

type BookRequestUpdateType struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	SubTitle    string      `json:"sub_title"`
	Description string      `json:"description"`
	Price       json.Number `json:"price"`
	Discount    int         `json:"discount" binding:"number"`
	Rating      int         `json:"rating" binding:"number"`
}
