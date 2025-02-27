package usecases

import (
	"errors"
	"log"
	"simplewebservice/internal/author"
	"simplewebservice/models"
	"simplewebservice/repositories"
	"strings"
)

type authorUsecaseImpl struct {
	AuthorRepository repositories.AuthorRepository
}

func NewAuthorUsecaseImpl(authorRepository repositories.AuthorRepository) *authorUsecaseImpl {
	return &authorUsecaseImpl{AuthorRepository: authorRepository}
}
func (a *authorUsecaseImpl) CreateAuthor(in *models.AuthorRequest) error {
	var author author.InsertAuthorDto
	author.Name = in.Name
	err := a.AuthorRepository.InsertAuthorSQL(&author)
	if err != nil {
		return err
	}
	return nil

}
func (a *authorUsecaseImpl) FindAuthor(id int) (*models.AuthorResponse, error) {
	var author models.AuthorResponse
	row, err := a.AuthorRepository.GetAuthorSQL(id)
	if err != nil {
		return nil, err
	}
	author.Name = row.Name
	return &author, nil
}

func (a *authorUsecaseImpl) FindAllAuthor(param string) (*[]models.AuthorResponse, error) {
	var authors []models.AuthorResponse
	if strings.ToLower(param) != "asc" && strings.ToLower(param) != "desc" {
		return nil, errors.New("0")
	}
	log.Println(param)
	rows, err := a.AuthorRepository.GetListAuthorSQL(param)
	if err != nil {
		return nil, err
	}
	for _, row := range *rows {
		authors = append(authors, models.AuthorResponse{
			Id:   row.Id,
			Name: row.Name,
		})
	}
	return &authors, nil
}

func (a *authorUsecaseImpl) DeleteAuthor(id int) error {
	err := a.AuthorRepository.DeleteAuthorSQL(id)
	if err != nil {
		return err
	}
	return nil
}

func (a *authorUsecaseImpl) UpdateAuthor(id int, in *models.AuthorRequest) error {
	var dto author.InsertAuthorDto
	dto.Name = in.Name
	err := a.AuthorRepository.UpdateAuthorSQL(id, &dto)
	if err != nil {
		return err
	}
	return nil
}
