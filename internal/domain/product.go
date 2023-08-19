package domain

type Product struct {
	ID           int
	Name         string
	Description  string
	ImageURL     string
	Price        float64
	CategoryCode string
	Stock        int
}
