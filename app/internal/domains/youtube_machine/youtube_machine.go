package ytmachine

import (
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	"github.com/igorshmel/lic_auto_post/app/pkg/ddo"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
)

// YoutubeMachine --
type YoutubeMachine struct {
	log             logger.Logger
	cfg             config.Config
	itemsNextCursor string
	playlistID      string
	itemsInfo       []ItemsInfo
}

// ItemsInfo --
type ItemsInfo struct {
	title   string
	videoId string
}

// NewYoutubeMachine - инициализация домена YoutubeMachine
func NewYoutubeMachine(log logger.Logger, cfg config.Config) *YoutubeMachine {
	log = log.WithMethod("VkMachineDomain")
	return &YoutubeMachine{log: log, cfg: cfg}
}

// KeepNextCursor --
func (ths *YoutubeMachine) KeepNextCursor(cursor string) {
	ths.itemsNextCursor = cursor
}

// GetNextCursor --
func (ths *YoutubeMachine) GetNextCursor() string {
	return ths.itemsNextCursor
}

// KeepItems --
func (ths *YoutubeMachine) KeepItems(title, videoID string) {
	ths.itemsInfo = append(ths.itemsInfo, ItemsInfo{title: title, videoId: videoID})
}

// GetItems --
func (ths *YoutubeMachine) GetItems() []ddo.ResGetItems {
	var getItems []ddo.ResGetItems

	for _, i := range ths.itemsInfo {
		getItems = append(getItems, ddo.ResGetItems{
			VideoId: i.videoId,
			Title:   i.title,
		})
	}

	return getItems
}
