package productmanager

type ProductResponse struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ImageUrl     string  `json:"imageUrl"`
	Price        float64 `json:"price"`
	CategoryCode string  `json:"categoryCode"`
	Stock        int     `json:"stock"`
}

type ProductCategoryResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
