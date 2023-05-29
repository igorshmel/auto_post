package events

// VkWallUploadEvent - структура события для публикации картинки на стену группы в ВК
type VkWallUploadEvent struct {
	FileName string `json:"file_name"`
}
