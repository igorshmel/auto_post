package routes

import (
	"auto_post/app/internal/adapters/repository"
	"auto_post/app/internal/adapters/transport/rest"
	"auto_post/app/internal/domain"
	"auto_post/app/internal/usecase/api"
	"git.fintechru.org/dfa/dfa_lib/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func registerRoutes(dom *domain.Dom, log logger.Logger, g *gin.Engine, repo *repository.Repository) {
	apiGroup := g.Group("/api")
	v1 := apiGroup.Group("/v1")

	// Создание usecase
	initParseImageUseCase := api.NewInitParseImageUseCase(log, repo.GetPersister(), repo.GetExtractor(), dom.GetParseImager())

	// Создание обработчиков запросов
	initParseImageEndpoint := rest.NewInitParseImageEndpoint(initParseImageUseCase, log)

	// Регистрация обработчиков запросов
	v1.POST("/init/", initParseImageEndpoint.Execute)
}

// Module ..
var Module = fx.Options(fx.Invoke(registerRoutes))
