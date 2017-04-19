package models

type User struct {
	Id           uint32 `json:"id"`
	Username     string `json:"username"`
	MoneyBalance uint32 `json:"balance"`
}

type UserParams struct {
	Username     string `json:"username"`
	MoneyBalance uint32 `json:"balance"`
}

