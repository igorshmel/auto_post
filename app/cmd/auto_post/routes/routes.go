package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/transport/rest"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/transport/youtube"
	manager "github.com/igorshmel/lic_auto_post/app/internal/domains/manager"
	vkMachine "github.com/igorshmel/lic_auto_post/app/internal/domains/vk_machine"
	ytmachine "github.com/igorshmel/lic_auto_post/app/internal/domains/youtube_machine"
	"github.com/igorshmel/lic_auto_post/app/internal/usecase/api"
	"github.com/igorshmel/lic_auto_post/app/internal/usecase/youtube_usecase"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	"github.com/igorshmel/lic_auto_post/app/pkg/deo"
	"github.com/igorshmel/lic_auto_post/app/pkg/dto"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/constants"
	"github.com/nuttech/bell/v2"
	"go.uber.org/fx"
)

func registerRoutes(
	g *gin.Engine,
	cfg config.Config,
	log logger.Logger,
	repo *repository.Repository,
	bellEvent *bell.Events,
	managerDomain *manager.Domain,
	vkMachineDomain *vkMachine.Domain,
	youtubeClient *youtube.YouTubeClient,
	youtubeDomain *ytmachine.Domain,
) {
	ctx := context.Background()

	apiGroup := g.Group("/api")
	v1 := apiGroup.Group("/v1")

	// Создание usecase
	createRecordUseCase := api.NewCreateRecordUseCase(cfg, log, bellEvent, repo.GetPersister(), repo.GetExtractor(), managerDomain.GetManagerPort(), vkMachineDomain.GetVkMachinePorts())
	proxyRecordUseCase := api.NewProxyRecordUseCase(cfg, log, bellEvent, repo.GetPersister(), repo.GetExtractor(), managerDomain.GetManagerPort(), vkMachineDomain.GetVkMachinePorts())
	downloadImageUseCase := api.NewDownloadImageUseCase(log, bellEvent, repo.GetPersister(), repo.GetExtractor(), managerDomain.GetManagerPort())
	vkWallUploadUseCase := api.NewVKWallPostUseCase(log, bellEvent, repo.GetPersister(), repo.GetExtractor(), vkMachineDomain.GetVkMachinePorts())

	getPlayListUseCase := youtube_usecase.NewGetPlayListUseCase(cfg, log, bellEvent, repo.GetPersister(), repo.GetExtractor(), youtubeClient, youtubeDomain.GetYoutubeMachinePorts())
	getYouTubeNextCursorUseCase := youtube_usecase.NewGetNextCursorUseCase(cfg, log, bellEvent, repo.GetPersister(), repo.GetExtractor(), youtubeClient, youtubeDomain.GetYoutubeMachinePorts())

	filteredYoutubeItemsUseCase := youtube_usecase.NewFilteredYoutubeItemsUseCase(cfg, log, bellEvent, repo.GetPersister(), repo.GetExtractor(), youtubeDomain.GetYoutubeMachinePorts())
	saveNewYoutubeItemsUseCase := youtube_usecase.NewSaveNewYoutubeItemsUseCase(cfg, log, bellEvent, repo.GetPersister(), repo.GetExtractor(), youtubeDomain.GetYoutubeMachinePorts())

	getRandomItemUseCase := youtube_usecase.NewGetRandomItemUseCase(cfg, log, bellEvent, repo.GetPersister(), repo.GetExtractor(), youtubeClient, youtubeDomain.GetYoutubeMachinePorts())

	// Создание обработчиков запросов
	createRecordEndpoint := rest.NewCreateRecordEndpoint(createRecordUseCase, log)
	proxyEndpoint := rest.NewProxyEndpoint(proxyRecordUseCase, log)
	downloadImageEndpoint := rest.NewDownloadImageEndpoint(downloadImageUseCase, log)

	// Регистрация обработчиков запросов
	v1.POST("/init/", createRecordEndpoint.CreateRecordExecute)
	v1.POST("/proxy/", proxyEndpoint.ProxyRecordExecute)
	v1.POST("/download/", downloadImageEndpoint.DownloadExecute)

	// add listener on event
	bellEvent.Listen(constants.DownloadImageEventName, func(msg bell.Message) {
		downloadImageEvent := msg.(deo.DownloadImageEvent)
		if err := downloadImageUseCase.Execute(nil, &dto.DownloadImageReqDTO{
			URL:    downloadImageEvent.Link,
			Output: downloadImageEvent.Output,
		}); err != nil {
			log.Error("failed execute DownloadImageInside with error: %s ", err.Error())
		}
		log.Info("DownloadImageEvent: %v", downloadImageEvent)
	})

	// прослушивание события vk_wall_upload
	bellEvent.Listen(constants.VkWallUploadEventName, func(msg bell.Message) {

		if err := vkWallUploadUseCase.Execute(nil); err != nil {
			log.Error("startEvent VkWallUploadEvent fail with error: %s ", err.Error())
		}
	})

	//////////////////////////////////////////////////
	// YOUTUBE
	//////////////////////////////////////////////////

	// get_next_cursor_youtube - событие
	bellEvent.Listen(constants.GetNextCursor_Youtube_Event, func(msg bell.Message) {

		// Добавляем requestID в контекст
		ctxWithRequestID, requestID := addRequestIDToContext(ctx)
		log.Info("Generated requestID: %s", requestID)

		log.Info("Event GetNextCursor")

		if err := getYouTubeNextCursorUseCase.Execute(ctxWithRequestID); err != nil {
			log.Error("Usecase error - GetNextCursorYoutubeUsecase with error: %s ", err.Error())
		}
	})

	// get_playlist_youtube - событие
	bellEvent.Listen(constants.GetPlaylist_Youtube_Event, func(msg bell.Message) {
		ctxEvent := msg.(context.Context)
		log.Info("Event GetPlayList")

		if err := getPlayListUseCase.Execute(ctxEvent); err != nil {
			log.Error("Usecase error - GetPlaylistYouTubeUsecase fail with error: %s ", err.Error())
		}
	})

	// youtube_get_playlist_done
	bellEvent.Listen(constants.Done_GetPlaylist_Event, func(msg bell.Message) {
		ctxEvent := msg.(context.Context)
		log.Info("Event DonePlayList")

		if err := filteredYoutubeItemsUseCase.Execute(ctxEvent); err != nil {
			log.Error("Usecase error - SaveNewYoutubeItemsUsecase with error: %s ", err.Error())
		}
	})

	// youtube_filtered_items_done
	bellEvent.Listen(constants.Done_Filtered_Items_Event, func(msg bell.Message) {
		ctxEvent := msg.(context.Context)
		log.Info("Event DonePlayList")

		if err := saveNewYoutubeItemsUseCase.Execute(ctxEvent); err != nil {
			log.Error("Usecase error - SaveNewYoutubeItemsUsecase with error: %s ", err.Error())
		}
	})

	//////////////////////////////////////////////////
	// VK
	//////////////////////////////////////////////////

	// get_video_vk - событие
	bellEvent.Listen(constants.Get_Random_Item_Event, func(msg bell.Message) {

		// Добавляем requestID в контекст
		ctxWithRequestID, requestID := addRequestIDToContext(ctx)
		log.Info("Generated requestID: %s", requestID)

		log.Info("Event GetRandomItem")

		if err := getRandomItemUseCase.Execute(ctxWithRequestID); err != nil {
			log.Error("Usecase error - GetRandomItemUseCase with error: %s ", err.Error())
		}
	})

}

// addRequestIDToContext -- создает requestID и добавляет его в контекст
func addRequestIDToContext(ctx context.Context) (context.Context, string) {
	requestID := uuid.New().String() // Создаем UUIDv4
	newCtx := context.WithValue(ctx, vars.RequestIDKey, requestID)
	return newCtx, requestID
}

// Module ..
var Module = fx.Options(fx.Invoke(registerRoutes))
