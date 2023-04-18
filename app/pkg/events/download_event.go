package events

// DownloadImageEvent - структура события для скачивания изображения
type DownloadImageEvent struct {
	Link   string `json:"link"`
	Output string `json:"output"`
}

// NewDownloadImageEvent is constructor
func NewDownloadImageEvent() *DownloadImageEvent {
	return &DownloadImageEvent{}
}
