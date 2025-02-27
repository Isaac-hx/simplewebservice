package usecases

import "simplewebservice/models"

type AuthorUsecase interface {
	CreateAuthor(in *models.AuthorRequest) error
	FindAuthor(id int) (*models.AuthorResponse, error)
	DeleteAuthor(id int) error
	UpdateAuthor(id int, in *models.AuthorRequest) error
	FindAllAuthor(param string) (*[]models.AuthorResponse, error)
}
