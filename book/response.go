package book

type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"Description"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
}