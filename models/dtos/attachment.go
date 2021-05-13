package dtos

type CreateAttachmentInput struct {
	FilePath string `json:"filepath" binding:"required"`
	FileType string `json:"filetype" binding:"required"`
}
