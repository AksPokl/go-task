package model

type User struct {
	Username string `json:"username" bson:"username,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
}

type UserResponse struct {
	Username string `json:"username"`
}

type TokenResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
