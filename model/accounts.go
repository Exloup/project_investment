package model

// Struct Customer Account
type CustomerAccounts struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone_number"`
	CreateAt string `json:"create_at"`
	Active   string `json:"is_activated"`
}
