package dto

type ProductResponse struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ImageURL     string  `json:"image_url"`
	Price        float64 `json:"price"`
	CategoryCode string  `json:"category_code"`
	Stock        int     `json:"stock"`
}

type ProductCategoryResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
