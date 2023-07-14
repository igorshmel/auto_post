package deo

// data event object

// DownloadImageEvent - структура события для скачивания изображения
type DownloadImageEvent struct {
	Link   string `json:"link"`
	Output string `json:"output"`
}
