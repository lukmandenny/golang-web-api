package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookRequestUpdate BookRequestUpdate) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err

	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	book := Book{
		Title:       bookRequest.Title,
		Price:       bookRequest.Price,
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
		Discount:    bookRequest.Discount,
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequestupdate BookRequestUpdate) (Book, error) {
	book, _ := s.repository.FindByID(ID)

	book.Title = bookRequestupdate.Title
	book.Price = bookRequestupdate.Price
	book.Description = bookRequestupdate.Description
	book.Rating = bookRequestupdate.Rating
	book.Discount = bookRequestupdate.Discount

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, _ := s.repository.FindByID(ID)

	newBook, err := s.repository.Delete(book)
	return newBook, err
}
