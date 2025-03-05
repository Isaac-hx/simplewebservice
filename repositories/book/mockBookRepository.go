package repositories

import "simplewebservice/book"

type MockBookPostgresRepository struct {
	Err  error
	Book *book.GetBookDto
}

func (m *MockBookPostgresRepository) GetBookSQL(id int) (*book.GetBookDto, error) {
	return m.Book, m.Err
}
