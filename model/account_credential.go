package model

// Struct Authen Account
type UserAccountCredential struct {
	Email    string `json:"userEmail"`
	Password string `json:"userPassword"`
	Name     string
}
