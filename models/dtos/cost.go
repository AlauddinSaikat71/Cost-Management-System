package dtos

import "time"

type CreateCostInput struct {
	Title       string    `json:"title" `
	Description string    `json:"description"`
	Amount      float32   `json:"amount" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	Payment_Id  int       `json:"payment_id"`
}

type UpdateCostInput struct {
	Title       string  `json:"title" `
	Description string  `json:"description"`
	Amount      float32 `json:"amount" binding:"required"`
	Payment_Id  int     `json:"payment_id"`
}
