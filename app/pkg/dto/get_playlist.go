package dto

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// NewGetPlayListReq is constructor
func NewGetPlayListReq() *GetPlayListReqDTO {
	return &GetPlayListReqDTO{}
}

// GetPlayListReqDTO --
type GetPlayListReqDTO struct {
	PlayListID string `json:"play_list_id"`
}

// Parse parses and validates the request
func (ths *GetPlayListReqDTO) Parse(c *gin.Context) error {
	return c.ShouldBindJSON(&ths)
}

// Validate validates an input request
func (ths *GetPlayListReqDTO) Validate() error {
	return validation.ValidateStruct(ths,
		validation.Field(&ths.PlayListID, validation.Required.Error("is required")),
	)
}

// GetPlayListResDTO --
type GetPlayListResDTO struct {
	Title   string `json:"title,omitempty"`
	VideoId string `json:"videoId,omitempty"`
}
