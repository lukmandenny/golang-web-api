package book

type BookRequestUpdate struct {
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Rating      int    `json:"rating" binding:"number"`
	Discount    int    `json:"discount" binding:"number"`
}
