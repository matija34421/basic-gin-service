package model

type Client struct {
	ID               int    `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	ResidenceAddress string `json:"residence_address"`
	BirthDate        string `json:"birth_date"`
}
