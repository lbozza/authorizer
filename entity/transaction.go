package entity

import "time"

type Transaction struct {
	ID       string    `json:"id,omitempty"`
	Merchant string    `json:"merchant"`
	Amount   int64     `json:"amount"`
	Time     time.Time `json:"time"`
}
