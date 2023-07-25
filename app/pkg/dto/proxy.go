package dto

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// NewProxyRecordReq is constructor
func NewProxyRecordReq() *ProxyRecordReqDTO {
	return &ProxyRecordReqDTO{}
}

// ProxyRecordReqDTO --
type ProxyRecordReqDTO struct {
	URL     string `json:"url"`
	AuthURL string `json:"auth_url"`
	Service string `json:"service"`
}

// Parse parses and validates the request
func (ths *ProxyRecordReqDTO) Parse(c *gin.Context) error {
	return c.ShouldBindJSON(&ths)
}

// Validate validates an input request
func (ths *ProxyRecordReqDTO) Validate() error {
	return validation.ValidateStruct(ths,
		validation.Field(&ths.URL, validation.Required.Error("is required"), is.URL),
		validation.Field(&ths.AuthURL, validation.Required.Error("is required")),
		validation.Field(&ths.Service, validation.Required.Error("is required")),
	)
}
