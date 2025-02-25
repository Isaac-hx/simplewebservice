package usecases

import (
	"simplewebservice/internal/book"
	"simplewebservice/models"
	"simplewebservice/repositories"
	"simplewebservice/utils"
)

type bookUsecaseImpl struct {
	BookRepository repositories.BookRepository
}

func NewBookUsecaseImpl(bookRepository repositories.BookRepository) *bookUsecaseImpl {
	return &bookUsecaseImpl{
		BookRepository: bookRepository,
	}
}

func (u *bookUsecaseImpl) CreateBook(in *models.BookRequest) error {
	published_date, err := utils.ParseTimeDate(in.PublishedDate)
	if err != nil {
		return err
	}

	dto := &book.InsertBookDto{
		Title:         in.Title,
		Description:   in.Description,
		TotalPage:     in.TotalPage,
		AuthorId:      in.AuthorId,
		PublishedDate: published_date,
		Price:         in.Price,
		CoverUrl:      in.CoverUrl,
	}
	//Perform logic business here\

	err = u.BookRepository.InsertBookSQL(dto)
	if err != nil {
		return err
	}

	return nil
}

func (u *bookUsecaseImpl) FindBook(id int) (*models.BookResponse, error) {
	var book models.BookResponse
	err := u.BookRepository.GetBookSQL(id).Scan(&book.Id, &book.Title, &book.AuthorId, &book.TotalPage, &book.Description, &book.PublishedDate, &book.Price, &book.CoverUrl)
	if err != nil {
		return nil, err
	}
	return &book, err
}
