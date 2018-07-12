package userModel

const UserCollection string = "users"

type UserSingup struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"email"`
}

type UserSignin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
