package dtos

type CreatePaymentInput struct {
	Method    string  `json:"method"`
	Amount    float32 `json:"amount" binding:"required"`
	CreatedBy int     `json:"createdby"`
	PaidBy    int     `json:"paidby"`
	Meta      string  `json:"meta"`
}

type UpdatePaymentInput struct {
	Method    string  `json:"method"`
	Amount    float32 `json:"amount"`
	CreatedBy int     `json:"createdby"`
	PaidBy    int     `json:"paidby"`
	Meta      string  `json:"meta"`
}
