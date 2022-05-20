package users

type User struct {
	Id       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}
