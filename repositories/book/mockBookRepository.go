// this file contains mock object implementation of bookRepository
package repositories

import (
	"simplewebservice/book"

	"github.com/stretchr/testify/mock"
)

type MockRepositoryBook struct {
	mock.Mock
}

// Make sure the implementation of bookRepository is completed
// this declaration is to verified the implementation object
var _ BookRepository = (*MockRepositoryBook)(nil)

// Implementatiion depedency repository book
func (m *MockRepositoryBook) InsertBookSQL(in *book.InsertBookDto) error {
	//Flag is the method has been called or no
	//Method called is used to return objects argument that contain value which is had been rule before.
	args := m.Called(in)
	//this method take the the first argument from object mocks.argument
	//this element is rule by you in test function mock.on().return()
	return args.Error(0)

}

func (m *MockRepositoryBook) GetBookSQL(id int) (*book.GetBookDto, error) {
	args := m.Called(id)
	return nil, args.Error(0)
}

func (m *MockRepositoryBook) GetListBookSQL(param string) (*[]book.GetBookDto, error) {
	args := m.Called(param)
	return nil, args.Error(0)
}

func (m *MockRepositoryBook) DeleteBookSQL(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRepositoryBook) UpdateBookSQL(id int, in *book.InsertBookDto) error {
	args := m.Called(id, in)
	return args.Error(0)
}
