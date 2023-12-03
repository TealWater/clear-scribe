package model

import "mime/multipart"

type IncomingText struct {
	Message string `json:"message"`
}

type FileUpload struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}
