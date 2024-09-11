package port

import (
	"github.com/igorshmel/lic_auto_post/app/pkg/dto"
)

// ChannelRepository --
type ChannelRepository interface {
	FindByID(id string) (*dto.Channel, error)
}

// CommentRepository --
type CommentRepository interface {
	FindByVideoID(videoID string) ([]dto.Comment, error)
}

// VideoRepository --
type VideoRepository interface {
	FindByID(id string) (*dto.Video, error)
}
