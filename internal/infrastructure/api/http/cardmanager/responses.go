package cardmanager

type Response struct {
	Type       string `json:"type"`
	Company    int    `json:"company"`
	HolderName string `json:"holderName"`
}
