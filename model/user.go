package model

type User struct {
	Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string
}

type CreateUserRequest struct {
	User
	StoreName string `json:"storeName`
	Password  string `json:"password"`
}

type UserToken struct {
	Model
	UserId      uint   `json:"userId"`
	AccessToken string `json:"accessToken"`
}
