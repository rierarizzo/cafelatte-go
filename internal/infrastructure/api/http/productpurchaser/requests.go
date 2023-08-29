package productpurchaser

type purchasedProduct struct {
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type CreateOrderRequest struct {
	UserId            int                `json:"userId"`
	AddressId         int                `json:"addressId"`
	PaymentMethodId   int                `json:"paymentMethodId"`
	Notes             string             `json:"notes"`
	PurchasedProducts []purchasedProduct `json:"purchasedProducts"`
}
