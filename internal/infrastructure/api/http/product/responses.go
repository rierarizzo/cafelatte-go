package product

type Response struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ImageURL     string  `json:"imageUrl"`
	Price        float64 `json:"price"`
	CategoryCode string  `json:"categoryCode"`
	Stock        int     `json:"stock"`
}

type CategoryResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
