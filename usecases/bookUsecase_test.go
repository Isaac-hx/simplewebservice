package usecases

import (
	"html"
	"simplewebservice/book"
	"simplewebservice/models"
	repositories "simplewebservice/repositories/book"
	"simplewebservice/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repo = repositories.MockRepositoryBook{mock.Mock{}}
var testUsecaseBook = NewBookUsecaseImpl(&repo)

func TestInsertBook_Success(t *testing.T) {
	dataBook := models.BookRequest{
		AuthorId:      1,
		Title:         "Sapiens: Riwayat Singkat Umat Manusia",
		Description:   "Narasi revolusioner tentang penciptaan dan evolusi umat manusia",
		TotalPage:     544,
		PublishedDate: "2014-09-04 00:00:00",
		Price:         200000,
		CoverUrl:      "https://m.media-amazon.com/images/I/713jIoMO3UL.jpg",
	}

	//setup mock
	//succees mock
	//if this method had been calling,the mock will return nil value
	publishedDate, _ := utils.ParseDate(dataBook.PublishedDate) // Gunakan fungsi parse
	expectedDto := &book.InsertBookDto{
		Title:         html.EscapeString(dataBook.Title),
		Description:   html.EscapeString(dataBook.Description),
		TotalPage:     dataBook.TotalPage,
		AuthorId:      dataBook.AuthorId,
		PublishedDate: *publishedDate,
		Price:         dataBook.Price,
		CoverUrl:      html.EscapeString(dataBook.CoverUrl),
	}
	t.Run("Success", func(t *testing.T) {
		repo.Mock.On("InsertBookSQL", expectedDto).Return(nil)
		err := testUsecaseBook.CreateBook(&dataBook)
		assert.NoError(t, err)
		repo.Mock.AssertExpectations(t)
	})
	t.Run("Invalid published date", func(t *testing.T) {
		invalidBook := models.BookRequest{
			AuthorId:      1,
			Title:         "Sapiens: Riwayat Singkat Umat Manusia",
			Description:   "Narasi revolusioner tentang penciptaan dan evolusi umat manusia",
			TotalPage:     544,
			PublishedDate: "invalid-date",
			Price:         200000,
			CoverUrl:      "https://m.media-amazon.com/images/I/713jIoMO3UL.jpg",
		}
		testcase := testUsecaseBook.CreateBook(&invalidBook)
		assert.Error(t, testcase)
		assert.Contains(t, testcase.Error(), "invalid date type")
		repo.Mock.AssertExpectations(t)

	})
}

func TestInsertBook_Failed(t *testing.T) {

}
