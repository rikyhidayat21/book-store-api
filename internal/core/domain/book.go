package domain

import (
	"github.com/rikyhidayat21/book-store-api/dto/bookDto"
	"strconv"
)

type Book struct {
	Id            int64   `db:"id" json:"id"`
	Title         string  `db:"title" json:"title"`
	YearPublished int64   `db:"year_published" json:"year_published"`
	Isbn          string  `db:"isbn" json:"isbn"`
	Price         float64 `db:"price" json:"price"`
	OutOfPrint    bool    `db:"out_of_print" json:"out_of_print"`
	Views         int64   `db:"views" json:"views"`
}

func (b Book) ToDto() bookDto.BookResponse {
	return bookDto.BookResponse{
		Id:            b.Id,
		Title:         b.Title,
		YearPublished: b.YearPublished,
		Isbn:          b.Isbn,
		Price:         b.Price,
		OutOfPrint:    b.OutOfPrint,
		Views:         b.Views,
	}
}

func (b Book) ToNewBookResponseDto() bookDto.NewBookResponse {
	return bookDto.NewBookResponse{Id: strconv.FormatInt(b.Id, 10)}
}
