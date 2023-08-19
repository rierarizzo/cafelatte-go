package productpurchaser

type purchasedProduct struct {
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
}

type CreateOrderRequest struct {
	UserID            int                `json:"userID"`
	AddressID         int                `json:"addressID"`
	PaymentMethodID   int                `json:"paymentMethodID"`
	Notes             string             `json:"notes"`
	PurchasedProducts []purchasedProduct `json:"purchasedProducts"`
}
