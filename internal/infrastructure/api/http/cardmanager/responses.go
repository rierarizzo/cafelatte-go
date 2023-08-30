package cardmanager

type CardResponse struct {
	Type       string `json:"type"`
	Company    int    `json:"company"`
	HolderName string `json:"holderName"`
}
