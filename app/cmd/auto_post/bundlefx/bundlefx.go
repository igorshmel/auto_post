package bundlefx

import (
	"auto_post/app/cmd/auto_post/configs"
	"auto_post/app/cmd/auto_post/domain"
	"auto_post/app/cmd/auto_post/evts"
	"auto_post/app/cmd/auto_post/ginserver"
	"auto_post/app/cmd/auto_post/log"
	"auto_post/app/cmd/auto_post/middleware"
	"auto_post/app/cmd/auto_post/repo"
	"auto_post/app/cmd/auto_post/routes"
	"auto_post/app/pkg/config"
	"context"
	"fmt"
	"git.fintechru.org/dfa/dfa_lib/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Module собирает все зависимости для запуска приложения
var Module = fx.Options(
	ginserver.Module,
	log.Module,
	configs.Module,
	domain.Module,
	routes.Module,
	repo.Module,
	evts.Module,

	fx.Invoke(setGinMiddlewares),
	fx.Invoke(setGinLogger),
	fx.Invoke(manageServer),
)

// manageServer управляет запуском и остановкой сервера
func manageServer(lc fx.Lifecycle, log logger.Logger, g *gin.Engine, cfg config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			port := cfg.App.Port
			log.Info("Starting application in :%s", port)

			go func() {
				err := g.Run(":" + port)
				if err != nil {
					panic(err)
				}
			}()
			return nil
		},

		OnStop: func(ctx context.Context) error {
			log.Info("Stopping application")
			return nil
		},
	})
}

// setGinLogger sets standard logger
func setGinLogger(router *gin.Engine) {
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		reqID := param.Request.Header.Get(middleware.RequestIDKey)
		if len(reqID) == 0 {
			reqID = "empty request id"
		}

		return fmt.Sprintf("%s [GIN]  [%v] |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
			param.TimeStamp.Format("2006/01/02 15:04:05"),
			reqID,
			param.StatusCodeColor(), param.StatusCode, param.ResetColor(),
			param.Latency,
			param.ClientIP,
			param.MethodColor(), param.Method, param.ResetColor(),
			param.Path,
			param.ErrorMessage,
		)
	}))
}

// setGinMiddlewares sets middlewares
func setGinMiddlewares(router *gin.Engine) {
	// middleware CORS settings
	router.Use(cors.Default())
	// middleware with request id for each request
	router.Use(middleware.MakeRequestIDGinMiddleware())
	// recovery middleware
	router.Use(gin.Recovery())
}
