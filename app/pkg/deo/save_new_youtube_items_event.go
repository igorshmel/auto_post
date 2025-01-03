package deo

// SaveNewYoutubeItemsEvent --
type SaveNewYoutubeItemsEvent struct {
	VideosInfo           []VideoInfo
	NextCursorPagination string
}

// VideoInfo --
type VideoInfo struct {
	Title   string `json:"title,omitempty"`
	VideoID string `json:"videoId,omitempty"`
}
