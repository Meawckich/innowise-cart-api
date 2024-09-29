package model

type Cart struct {
	Id    int        `db:"id"`
	Items []CartItem `json:"items"`
}
