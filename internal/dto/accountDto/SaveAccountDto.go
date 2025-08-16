package accountdto

type SaveAccountDto struct {
	ClientId      int    `json:"client_id"`
	AccountNumber string `json:"account_number"`
	Balance       int    `json:"balance"`
}
