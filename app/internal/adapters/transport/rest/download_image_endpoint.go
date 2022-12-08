package rest

import (
	"net/http"

	"auto_post/app/cmd/auto_post/middleware"
	"auto_post/app/internal/adapters/port"
	"auto_post/app/pkg/dto"
	"git.fintechru.org/dfa/dfa_lib/errs"
	"git.fintechru.org/dfa/dfa_lib/logger"
	"github.com/gin-gonic/gin"
)

// DownloadImageEndpoint --
type DownloadImageEndpoint struct {
	log     logger.Logger
	usecase port.DownloadImageUseCase
}

// NewDownloadImageEndpoint _
func NewDownloadImageEndpoint(usecase port.DownloadImageUseCase, log logger.Logger) port.Endpoint {
	return DownloadImageEndpoint{
		log:     log,
		usecase: usecase,
	}
}

// Execute is handler
func (ths DownloadImageEndpoint) Execute(ctx *gin.Context) {
	ths.log = middleware.SetRequestIDPrefix(ctx, ths.log)
	log := ths.log.WithMethod("endpoint CreateInvestor")

	req := dto.NewReqDownloadImage()

	// request parse
	if err := req.Parse(ctx); err != nil {
		log.Error("unable to parse a request: %s", err)
		ctx.JSON(http.StatusOK, errs.New().SetCode(errs.ParseRequest).
			SetMsg("unable to parse a request").
			GinJSON())
		return
	}

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
		log.Error("failed call to usecase CreateInvestor: %s", err)
		ctx.JSON(http.StatusOK, errs.FromError(err).GinJSON())
		return
	}

	ctx.Status(http.StatusOK)
}
