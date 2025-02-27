package usecases

import (
	"errors"
	"simplewebservice/internal/book"
	"simplewebservice/models"
	"simplewebservice/repositories"
	"simplewebservice/utils"
	"strings"
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
	row, err := u.BookRepository.GetBookSQL(id)
	if err != nil {
		return nil, err
	}
	book = models.BookResponse{Id: row.Id, Title: row.Title, AuthorName: row.AuthorName, Description: row.Description, TotalPage: row.TotalPage, CoverUrl: row.CoverUrl, Price: row.Price, PublishedDate: row.PublishedDate}
	return &book, err
}

func (u *bookUsecaseImpl) DeleteBook(id int) error {
	err := u.BookRepository.DeleteBookSQL(id)
	if err != nil {
		return err
	}
	return nil

}
func (u *bookUsecaseImpl) EditBook(id int, in *models.BookRequest) error {
	published_date, err := utils.ParseTimeDate(in.PublishedDate)
	if err != nil {
		return err
	}
	dto := &book.InsertBookDto{
		Title:         in.Title,
		Description:   in.Description,
		AuthorId:      in.AuthorId,
		TotalPage:     in.TotalPage,
		PublishedDate: published_date,
		Price:         in.Price,
		CoverUrl:      in.CoverUrl,
	}
	err = u.BookRepository.UpdateBookSQL(id, dto)
	if err != nil {
		return err
	}
	return nil
}

func (u *bookUsecaseImpl) FindAllBook(param string) (*[]models.BookResponse, error) {
	var books []models.BookResponse
	if strings.ToLower(param) != "asc" && strings.ToLower(param) != "desc" {
		return nil, errors.New("error : invalid query param")
	}
	rows, err := u.BookRepository.GetListBookSQL(param)
	if err != nil {
		return nil, err
	}
	for _, eachBook := range *rows {
		books = append(books, models.BookResponse{
			Id:         eachBook.Id,
			Price:      float64(eachBook.Price),
			Title:      eachBook.Title,
			AuthorName: eachBook.AuthorName,
			CoverUrl:   eachBook.CoverUrl,
		})

	}
	return &books, nil
}
