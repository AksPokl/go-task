package model

type UserResponse struct {
	Username string `json:"username"`
}

type TokenResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
