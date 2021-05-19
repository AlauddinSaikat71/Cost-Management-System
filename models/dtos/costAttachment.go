package dtos

type CreateCostAttachmentInput struct {
	CostId       int `json:"costid"`
	AttachmentId int `json:"attachmentid"`
	CreatedBy    int `json:"createdby"`
}
