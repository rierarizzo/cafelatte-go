package purchase

type PurchasedProduct struct {
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
}

type OrderRequest struct {
	UserID            int                `json:"userID"`
	AddressID         int                `json:"addressID"`
	PaymentMethodID   int                `json:"paymentMethodID"`
	Notes             string             `json:"notes"`
	PurchasedProducts []PurchasedProduct `json:"purchasedProducts"`
}
