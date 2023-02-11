package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(bookData BookRequestType) (Book, error)
	Update(ID int, bookData BookRequestUpdateType) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func AssignService(db *repository) *service {
	return &service{db}
}

func (s *service) FindAll() ([]Book, error) {
	// books, err := s.repository.FindAll()
	// return books, err
	return s.repository.FindAll()
}

func (s *service) FindById(ID int) (Book, error) {
	return s.repository.FindById(ID)
}

func (s *service) Create(bookData BookRequestType) (Book, error) {
	price, _ := bookData.Price.Int64()

	book := Book{
		Title:       bookData.Title,
		SubTitle:    bookData.SubTitle,
		Description: bookData.Description,
		Price:       int(price),
		Rating:      bookData.Rating,
		Discount:    bookData.Discount,
	}
	return s.repository.Create(book)
}

func (s *service) Update(ID int, bookData BookRequestUpdateType) (Book, error) {
	book, err := s.repository.FindById(ID)
	price, _ := bookData.Price.Int64()

	if len(bookData.Title) > 0 {
		book.Title = bookData.Title
	}
	if len(bookData.SubTitle) > 0 {
		book.SubTitle = bookData.SubTitle
	}
	if len(bookData.Description) > 0 {
		book.Description = bookData.Description
	}
	if price > 0 {
		book.Price = int(price)
	}
	if bookData.Rating > 0 {
		book.Rating = bookData.Rating
	}
	if bookData.Discount > 0 {
		book.Discount = bookData.Discount
	}

	updatedBook, err := s.repository.Update(book)
	return updatedBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	newBook, err := s.repository.Delete(book)
	return newBook, err
}
