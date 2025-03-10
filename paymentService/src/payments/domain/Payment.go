package domain

type Payment struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}
