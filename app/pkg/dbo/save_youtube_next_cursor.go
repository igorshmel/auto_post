package dbo

import "time"

// YoutubeNextCursorDBO --
type YoutubeNextCursorDBO struct {
	NextPageToken string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}
