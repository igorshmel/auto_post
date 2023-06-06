package bundlefx

import (
	"auto_post/app/cmd/auto_post/chan_os"
	"auto_post/app/cmd/auto_post/configs"
	"auto_post/app/cmd/auto_post/cron"
	"auto_post/app/cmd/auto_post/domains"
	"auto_post/app/cmd/auto_post/events"
	"auto_post/app/cmd/auto_post/ginserver"
	"auto_post/app/cmd/auto_post/log"
	"auto_post/app/cmd/auto_post/middleware"
	"auto_post/app/cmd/auto_post/repo"
	"auto_post/app/cmd/auto_post/routes"
	"auto_post/app/pkg/config"
	"auto_post/app/pkg/deo"
	logger "auto_post/app/pkg/log"
	"auto_post/app/pkg/vars/constants"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/nuttech/bell/v2"
	"go.uber.org/fx"
)

// Module собирает все зависимости для запуска приложения
var Module = fx.Options(
	ginserver.Module,
	log.Module,
	configs.Module,
	routes.Module,
	repo.Module,
	events.Module,
	cron.Module,
	chan_os.Module,
	domains.ManagerDomainModule,
	domains.VkMachineDomainModule,

	fx.Invoke(setGinMiddlewares),
	fx.Invoke(setGinLogger),
	fx.Invoke(manageServer),
	fx.Invoke(setCron),
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

func setCron(
	log logger.Logger,
	cron *cron.Runner,
	bellEvent *bell.Events,
) {
	task := func(in string, job gocron.Job) {
		fmt.Printf("this job's last run: %s this job's next run: %s\n", job.LastRun(), job.NextRun())
		fmt.Printf("in argument is %s\n", in)

		// отправка события в домен VkMachineDomain
		if err := bellEvent.Ring(
			constants.VkWallUploadEventName, deo.VkWallUploadEvent{FileName: ""}); err != nil {
			log.Error("unable send event VkWallUpload with error: %s", err.Error())
		}

		log.Debug("sendEvent VkWallUploadEvent success!")
	}

	// Конфигурируем время и частоту выполнения задачи
	if _, err := cron.Cron("*/1 * * * *").DoWithJobDetails(task, "foo"); err != nil {
		log.Error("unable to set the task: %s", err)
		return
	}
	cron.StartAsync()
}
