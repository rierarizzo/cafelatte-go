package domain

type Product struct {
	Id           int
	Name         string
	Description  string
	ImageUrl     string
	Price        float64
	CategoryCode string
	Stock        int
}

type ProductCategory struct {
	Code        string
	Description string
}
