package book

type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	SubTitle    string `json:"sub_title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Discount    int    `json:"discount"`
	Rating      int    `json:"rating"`
}

func BooksResponse(book Book) BookResponse {
	return BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		SubTitle:    book.SubTitle,
		Description: book.Description,
		Price:       book.Price,
		Discount:    book.Discount,
		Rating:      book.Rating,
	}
}
