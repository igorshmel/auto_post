package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

import validation "github.com/go-ozzo/ozzo-validation/v4"

// CreateRecordReqDTO --
type CreateRecordReqDTO struct {
	URL         string `json:"url"`
	AuthURL     string `json:"auth_url"`
	ImgURL      string `json:"img_url"`
	Description string `json:"description"`
	Service     string `json:"service"`
}

// Parse parses and validates the request
func (ths *CreateRecordReqDTO) Parse(c *gin.Context) error {
	return c.ShouldBindJSON(&ths)
}

// Validate validates an input request
func (ths *CreateRecordReqDTO) Validate() error {
	return validation.ValidateStruct(ths,
		validation.Field(&ths.URL, validation.Required.Error("is required"), is.URL),
		validation.Field(&ths.Service, validation.Required.Error("is required")),
	)
}
