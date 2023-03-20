package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

import validation "github.com/go-ozzo/ozzo-validation/v4"

// DownloadImageReqDTO --
type DownloadImageReqDTO struct {
	URL    string `json:"url"`
	Output string `json:"output"`
}

// NewDownloadImageReq is constructor
func NewDownloadImageReq() *DownloadImageReqDTO {
	return &DownloadImageReqDTO{}
}

// Parse and validates the request
func (ths *DownloadImageReqDTO) Parse(c *gin.Context) error {
	return c.ShouldBindJSON(&ths)
}

// Validate validates an input request
func (ths *DownloadImageReqDTO) Validate() error {
	return validation.ValidateStruct(ths,
		validation.Field(&ths.URL, validation.Required.Error("is required"), is.URL),
		validation.Field(&ths.Output, validation.Required.Error("is required")),
	)
}
