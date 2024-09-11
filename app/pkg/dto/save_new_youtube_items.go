package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// NewSaveNewYoutubeItemsReq is constructor
func NewSaveNewYoutubeItemsReq() *SaveNewYoutubeItemsReqDTO {
	return &SaveNewYoutubeItemsReqDTO{}
}

// SaveNewYoutubeItemsReqDTO --
type SaveNewYoutubeItemsReqDTO struct {
	VideosInfo           []VideoInfo
	NextCursorPagination string
}

// VideoInfo --
type VideoInfo struct {
	Title   string `json:"title,omitempty"`
	VideoId string `json:"videoId,omitempty"`
}

// Validate validates an input request
func (ths *SaveNewYoutubeItemsReqDTO) Validate() error {
	return validation.ValidateStruct(ths,
		validation.Field(&ths.VideosInfo, validation.Required.Error("is required")),
	)
}
