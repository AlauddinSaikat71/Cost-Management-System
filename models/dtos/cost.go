package dtos

type CreateCostInput struct {
	Title       string  `json:"title" `
	Description string  `json:"description"`
	Amount      float32 `json:"amount" binding:"required"`
	Date        string  `json:"date" binding:"required"`
	Payment_Id  int     `json:"payment_id"`
}

type UpdateCostInput struct {
	Title       string  `json:"title" `
	Description string  `json:"description"`
	Amount      float32 `json:"amount" binding:"required"`
	Payment_Id  int     `json:"payment_id"`
	Date        string  `json:"date" binding:"required"`
}
