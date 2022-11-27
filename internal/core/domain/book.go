package domain

type Book struct {
	Id            string  `db:"id" json:"id"`
	Title         string  `db:"title" json:"title"`
	YearPublished int64   `db:"year_published" json:"year_published"`
	Isbn          string  `db:"isbn" json:"isbn"`
	Price         float64 `db:"price" json:"price"`
	OutOfPrint    bool    `db:"out_of_print" json:"out_of_print"`
	Views         int64   `db:"views" json:"views"`
}
