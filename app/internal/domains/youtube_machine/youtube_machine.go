package ytmachine

import (
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	"github.com/igorshmel/lic_auto_post/app/pkg/ddo"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"sync"
)

// YoutubeMachineState -- состояние отдельного сценария
type YoutubeMachineState struct {
	itemsNextCursor string
	itemsInfo       []ItemsInfo
	randomItem      RandomItem
}

// YoutubeMachine -- основной домен
type YoutubeMachine struct {
	log    logger.Logger
	cfg    config.Config
	states map[string]*YoutubeMachineState
	mu     sync.Mutex // мьютекс для безопасности потоков
}

// ItemsInfo -- структура для хранения информации об элементе
type ItemsInfo struct {
	title   string
	videoID string
}

// RandomItem --
type RandomItem struct {
	VideoID string
	Title   string
}

// NewYoutubeMachine -- инициализация домена YoutubeMachine
func NewYoutubeMachine(log logger.Logger, cfg config.Config) *YoutubeMachine {
	log = log.WithMethod("YoutubeMachineDomain")
	return &YoutubeMachine{
		log:    log,
		cfg:    cfg,
		states: make(map[string]*YoutubeMachineState),
	}
}

// GetOrCreateState -- получение или создание состояния для уникального сценария
func (ths *YoutubeMachine) GetOrCreateState(requestID string) *YoutubeMachineState {
	ths.mu.Lock()
	defer ths.mu.Unlock()

	if state, exists := ths.states[requestID]; exists {
		return state
	}

	state := &YoutubeMachineState{}
	ths.states[requestID] = state
	return state
}

// KeepNextCursor -- сохраняет курсор для указанного сценария
func (ths *YoutubeMachine) KeepNextCursor(requestID, cursor string) {
	state := ths.GetOrCreateState(requestID)
	state.itemsNextCursor = cursor
}

// KeepRandomItem -- сохраняет случайную запись о видеоролике
func (ths *YoutubeMachine) KeepRandomItem(requestID string, ddo *ddo.YoutubeItemDDO) {
	state := ths.GetOrCreateState(requestID)
	state.randomItem.Title = ddo.Title
	state.randomItem.VideoID = ddo.VideoID
}

// GetRandomItem -- возвращает случайную запись о видеоролике
func (ths *YoutubeMachine) GetRandomItem(requestID string) *ddo.YoutubeItemDDO {
	state := ths.GetOrCreateState(requestID)
	return &ddo.YoutubeItemDDO{
		Title:   state.randomItem.Title,
		VideoID: state.randomItem.VideoID,
	}
}

// GetNextCursor -- получает курсор для указанного сценария
func (ths *YoutubeMachine) GetNextCursor(requestID string) string {
	state := ths.GetOrCreateState(requestID)
	return state.itemsNextCursor
}

// KeepItems -- сохраняет элементы для указанного сценария
func (ths *YoutubeMachine) KeepItems(requestID, title, videoID string) {
	state := ths.GetOrCreateState(requestID)
	state.itemsInfo = append(state.itemsInfo, ItemsInfo{title: title, videoID: videoID})
}

// GetItems -- получает элементы для указанного сценария
func (ths *YoutubeMachine) GetItems(requestID string) []ddo.ResGetItems {
	state := ths.GetOrCreateState(requestID)
	var getItems []ddo.ResGetItems

	for _, i := range state.itemsInfo {
		getItems = append(getItems, ddo.ResGetItems{
			VideoID: i.videoID,
			Title:   i.title,
		})
	}

	return getItems
}

// DeleteState -- удаляет состояние для указанного сценария
func (ths *YoutubeMachine) DeleteState(requestID string) {
	ths.mu.Lock()
	defer ths.mu.Unlock()

	if _, exists := ths.states[requestID]; exists {
		delete(ths.states, requestID)
	}
}

// ClearItems -- очищает список элементов для указанного сценария
func (ths *YoutubeMachine) ClearItems(requestID string) {
	state := ths.GetOrCreateState(requestID)
	state.itemsInfo = []ItemsInfo{} // Сбрасываем список
}
