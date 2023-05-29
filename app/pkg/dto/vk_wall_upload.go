package dto

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// VkWallUploadReqDTO --
type VkWallUploadReqDTO struct {
	FileName string `json:"file_name"`
}

// NewVkWallUploadReqDTO is constructor
func NewVkWallUploadReqDTO() *VkWallUploadReqDTO {
	return &VkWallUploadReqDTO{}
}

// Parse and validates the request
func (ths *VkWallUploadReqDTO) Parse(c *gin.Context) error {
	return c.ShouldBindJSON(&ths)
}

// Validate validates an input request
func (ths *VkWallUploadReqDTO) Validate() error {
	return validation.ValidateStruct(ths) //validation.Field(&ths.FileName, validation.Required.Error("is required"), is.URL),

}
