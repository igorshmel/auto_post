package deo

// SaveNewYoutubeItemsEvent --
type SaveNewYoutubeItemsEvent struct {
	VideosInfo           []VideoInfo
	NextCursorPagination string
}

// VideoInfo --
type VideoInfo struct {
	Title   string `json:"title,omitempty"`
	VideoId string `json:"videoId,omitempty"`
}
