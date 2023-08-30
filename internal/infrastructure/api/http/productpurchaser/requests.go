package productpurchaser

type OrderCreate struct {
	UserId            int    `json:"userId"`
	AddressId         int    `json:"addressId"`
	PaymentMethodId   int    `json:"paymentMethodId"`
	Notes             string `json:"notes"`
	PurchasedProducts []struct {
		ProductId int `json:"productId"`
		Quantity  int `json:"quantity"`
	} `json:"purchasedProducts"`
}
