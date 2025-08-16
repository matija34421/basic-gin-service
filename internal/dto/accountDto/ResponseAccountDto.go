package accountdto

import "time"

type ResponseAccountDto struct {
	Id            int       `json:"id"`
	ClientId      int       `json:"client_id"`
	AccountNumber string    `json:"account_number"`
	Balance       int       `json:"balance"`
	Created_at    time.Time `json:"created_at"`
}
