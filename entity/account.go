package entity

type Account struct {
	ID             string `json:"id,omitempty"`
	ActiveCard     bool   `json:"active-card"`
	AvaliableLimit int64  `json:"available-limit"`
}
