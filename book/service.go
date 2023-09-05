package book

type Service interface {
	FindAll() ([]Book, error)
	Create(book BookInput) (Book, error)
	FindByID(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookinput BookInput) (Book, error) {
	price, _ := bookinput.Price.Int64()
	rating, _ := bookinput.Rating.Int64()

	book := Book{
		Title:       bookinput.Title,
		Price:       int(price),
		Description: bookinput.Description,
		Rating:      int(rating),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}