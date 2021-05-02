package dtos

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Name     string `json:"name" `
	Email    string `json:"email" `
	Phone    string `json:"phone" `
	Password string `json:"password" `
}
