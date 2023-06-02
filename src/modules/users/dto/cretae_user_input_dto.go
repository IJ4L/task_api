package users

type CreateUserInputDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
