package userModel

const UserCollection string = "users"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
