package model

import "time"

type Account struct {
	Id            int       `json:"id"`
	ClientId      int       `json:"client_id"`
	AccountNumber string    `json:"account_number"`
	Balance       float64   `json:"balance"`
	Created_at    time.Time `json:"created_at"`
}
