package routes

import (
	"auto_post/app/internal/adapters/repository"
	"auto_post/app/internal/adapters/transport/rest"
	"auto_post/app/internal/domain"
	"auto_post/app/internal/usecase/api"
	"auto_post/app/internal/usecase/inside"
	"auto_post/app/pkg/dto"
	"auto_post/app/pkg/events"
	logger "auto_post/app/pkg/log"
	"auto_post/app/pkg/vars/constants"
	"github.com/gin-gonic/gin"
	"github.com/nuttech/bell/v2"
	"go.uber.org/fx"
)

func registerRoutes(dom *domain.Dom, log logger.Logger, g *gin.Engine, repo *repository.Repository, bellEvent *bell.Events) {
	apiGroup := g.Group("/api")
	v1 := apiGroup.Group("/v1")

	// Создание usecase
	initParseImageUseCase := api.NewInitParseImageUseCase(log, repo.GetPersister(), repo.GetExtractor(), dom.GetParseImager(), bellEvent)
	downloadImageUseCase := api.NewDownloadImageUseCase(log, repo.GetPersister(), repo.GetExtractor(), dom.GetParseImager())

	// Создание обработчиков запросов
	initParseImageEndpoint := rest.NewInitParseImageEndpoint(initParseImageUseCase, log)
	downloadImageEndpoint := rest.NewDownloadImageEndpoint(downloadImageUseCase, log)

	// Регистрация обработчиков запросов
	v1.POST("/init/", initParseImageEndpoint.Execute)
	v1.POST("/download/", downloadImageEndpoint.Execute)
	downloadImageInside := inside.NewDownloadImageInside(log)

	// add listener on event
	bellEvent.Listen(constants.DownloadImageEventName, func(msg bell.Message) {
		downloadImageEvent := msg.(events.DownloadImageEvent)
		if err := downloadImageInside.Execute(&dto.DownloadImageReqDTO{
			URL:    downloadImageEvent.Link,
			Output: downloadImageEvent.Output,
		}); err != nil {
			log.Error("failed execute DownloadImageInside with error: %s ", err.Error())
		}
		log.Info("DownloadImageEvent: %v", downloadImageEvent)
	})
}

// Module ..
var Module = fx.Options(fx.Invoke(registerRoutes))
