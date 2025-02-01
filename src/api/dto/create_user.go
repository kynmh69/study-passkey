package dto

type CreateUserInput struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}
