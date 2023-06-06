package deo

// data event object

// VkWallUploadEvent - структура события для публикации картинки на стену группы в ВК
type VkWallUploadEvent struct {
	FileName string `json:"file_name"`
}
