package model

type User struct {
	Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string
}

type UserToken struct {
	Model
	UserId      uint   `json:"userId"`
	AccessToken string `json:"accessToken"`
}
