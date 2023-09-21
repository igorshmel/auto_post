package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/transport/rest"
	manager "github.com/igorshmel/lic_auto_post/app/internal/domains/manager"
	vkMachine "github.com/igorshmel/lic_auto_post/app/internal/domains/vk_machine"
	"github.com/igorshmel/lic_auto_post/app/internal/usecase/api"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	"github.com/igorshmel/lic_auto_post/app/pkg/deo"
	"github.com/igorshmel/lic_auto_post/app/pkg/dto"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
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
) {
	apiGroup := g.Group("/api")
	v1 := apiGroup.Group("/v1")

	// Создание usecase
	createRecordUseCase := api.NewCreateRecordUseCase(cfg, log, bellEvent, repo.GetPersister(), repo.GetExtractor(), managerDomain.GetManagerPort(), vkMachineDomain.GetVkMachinePorts())
	proxyRecordUseCase := api.NewProxyRecordUseCase(cfg, log, bellEvent, repo.GetPersister(), repo.GetExtractor(), managerDomain.GetManagerPort(), vkMachineDomain.GetVkMachinePorts())
	downloadImageUseCase := api.NewDownloadImageUseCase(log, bellEvent, repo.GetPersister(), repo.GetExtractor(), managerDomain.GetManagerPort())
	vkWallUploadUseCase := api.NewVKWallPostUseCase(log, bellEvent, repo.GetPersister(), repo.GetExtractor(), vkMachineDomain.GetVkMachinePorts())

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
}

// Module ..
var Module = fx.Options(fx.Invoke(registerRoutes))
