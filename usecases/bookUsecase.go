// This file contain depedency of instance or object usecase

package usecases

import (
	"simplewebservice/models"
)

type BookUsecase interface {
	CreateBook(in *models.BookRequest) error
	FindBook(id int) (*models.BookResponse, error)
	DeleteBook(id int) error
	EditBook(id int, in *models.BookRequest) error
	FindAllBook(param string) (*[]models.BookResponse, error)
}
