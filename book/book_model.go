package book

import "time"

type Book struct {
	ID          int
	Title       string
	SubTitle    string
	Description string
	Price       int
	Discount    int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
