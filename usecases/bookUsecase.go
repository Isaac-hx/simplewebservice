// This file contain depedency of instance or object usecase

package usecases

import (
	"simplewebservice/models"
)

type BookUsecase interface {
	CreateBook(in *models.BookRequest) error
	FindBook(id int) (*models.BookResponse, error)
}
