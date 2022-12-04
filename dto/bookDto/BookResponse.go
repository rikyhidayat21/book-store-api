package bookDto

type BookResponse struct {
	Id            string  `json:"id"`
	Title         string  `json:"title"`
	YearPublished int64   `json:"year_published"`
	Isbn          string  `json:"isbn"`
	Price         float64 `json:"price"`
	OutOfPrint    bool    `json:"out_of_print"`
	Views         int64   `json:"views"`
}
