package dto

type PurchasedProduct struct {
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
}

type PurchaseOrderRequest struct {
	UserID            int                `json:"userID"`
	AddressID         int                `json:"addressID"`
	PaymentMethodID   int                `json:"paymentMethodID"`
	Notes             string             `json:"notes"`
	PurchasedProducts []PurchasedProduct `json:"purchasedProducts"`
}
