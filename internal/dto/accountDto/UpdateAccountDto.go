package accountdto

type UpdateAccountDto struct {
	Id      int     `json:"id"`
	Amount  float64 `json:"amount"`
	Deposit bool    `json:"deposit"`
}
