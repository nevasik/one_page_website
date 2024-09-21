package model

type Order struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	Telegram    string `json:"telegram"`
}

type OrderResp struct {
}
