package routes

import (
	"auto_post/app/internal/adapters/repository"
	"auto_post/app/internal/adapters/transport/rest"
	manager "auto_post/app/internal/domains/manager"
	vkMachine "auto_post/app/internal/domains/manager"
	"auto_post/app/internal/usecase/api"
	"auto_post/app/pkg/config"
	"auto_post/app/pkg/dto"
	"auto_post/app/pkg/events"
	logger "auto_post/app/pkg/log"
	"auto_post/app/pkg/vars/constants"
	"github.com/gin-gonic/gin"
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
) {
	apiGroup := g.Group("/api")
	v1 := apiGroup.Group("/v1")

	// Создание usecase
	createRecordUseCase := api.NewCreateRecordUseCase(cfg, log, bellEvent, repo.GetPersister(), repo.GetExtractor(), managerDomain.GetManagerPort())
	downloadImageUseCase := api.NewDownloadImageUseCase(log, repo.GetPersister(), repo.GetExtractor(), managerDomain.GetManagerPort())

	// Создание обработчиков запросов
	createRecordEndpoint := rest.NewCreateRecordEndpoint(createRecordUseCase, log)
	downloadImageEndpoint := rest.NewDownloadImageEndpoint(downloadImageUseCase, log)

	// Регистрация обработчиков запросов
	v1.POST("/init/", createRecordEndpoint.CreateRecordExecute)
	v1.POST("/download/", downloadImageEndpoint.DownloadExecute)

	// add listener on event
	bellEvent.Listen(constants.DownloadImageEventName, func(msg bell.Message) {
		downloadImageEvent := msg.(events.DownloadImageEvent)
		if err := downloadImageUseCase.Execute(nil, &dto.DownloadImageReqDTO{
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
