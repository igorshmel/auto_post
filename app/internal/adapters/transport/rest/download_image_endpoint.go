package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igorshmel/lic_auto_post/app/cmd/auto_post/middleware"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/port"
	"github.com/igorshmel/lic_auto_post/app/pkg/dto"
	"github.com/igorshmel/lic_auto_post/app/pkg/errs"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
)

// DownloadImageEndpoint --
type DownloadImageEndpoint struct {
	log     logger.Logger
	usecase port.DownloadImageUseCase
}

// NewDownloadImageEndpoint --
func NewDownloadImageEndpoint(usecase port.DownloadImageUseCase, log logger.Logger) port.DownloadEndpoint {
	return DownloadImageEndpoint{
		log:     log,
		usecase: usecase,
	}
}

// DownloadExecute is handler
func (ths DownloadImageEndpoint) DownloadExecute(ctx *gin.Context) {
	ths.log = middleware.SetRequestIDPrefix(ctx, ths.log)
	log := ths.log.WithMethod("endpoint CreateRecord")

	req := dto.NewDownloadImageReq()

	// request parse
	if err := req.Parse(ctx); err != nil {
		log.Error("unable to download a request: %s", err)
		ctx.JSON(http.StatusOK, errs.New().SetCode(errs.ParseRequest).
			SetMsg("unable to download a request").
			GinJSON())
		return
	}

	log.Debug("req: %v", *req)

	// validate request
	if err := req.Validate(); err != nil {
		log.Error("error of validation: %s", err)
		ctx.JSON(http.StatusOK, errs.New().SetCode(errs.Syntax).
			SetMsg(err.Error()).
			GinJSON())
		return
	}

	err := ths.usecase.Execute(ctx, req)
	if err != nil {
		log.Error("failed call to usecase DownloadImage: %s", err)
		ctx.JSON(http.StatusOK, errs.FromError(err).GinJSON())
		return
	}

	ctx.Status(http.StatusOK)
}
