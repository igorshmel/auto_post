package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/igorshmel/lic_auto_post/app/cmd/auto_post/middleware"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/port"
	"github.com/igorshmel/lic_auto_post/app/pkg/dto"
	"github.com/igorshmel/lic_auto_post/app/pkg/errs"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"net/http"
)

// ProxyRecordEndpoint --
type ProxyRecordEndpoint struct {
	log     logger.Logger
	usecase port.ProxyRecordUseCase
}

// NewProxyEndpoint --
func NewProxyEndpoint(usecase port.ProxyRecordUseCase, log logger.Logger) port.ProxyEndpoint {
	return ProxyRecordEndpoint{
		log:     log,
		usecase: usecase,
	}
}

// ProxyRecordExecute is handler
func (ths ProxyRecordEndpoint) ProxyRecordExecute(ctx *gin.Context) {
	ths.log = middleware.SetRequestIDPrefix(ctx, ths.log)
	log := ths.log.WithMethod("endpoint ProxyRecord")

	req := dto.NewProxyRecordReq()

	// request parse
	if err := req.Parse(ctx); err != nil {
		log.Error("unable to proxy record a request: %s", err)
		ctx.JSON(http.StatusOK, errs.New().SetCode(errs.ParseRequest).
			SetMsg("unable to proxy record a request").
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
		log.Error("failed call to usecase ProxyRecord: %s", err)
		ctx.JSON(http.StatusOK, errs.FromError(err).GinJSON())
		return
	}

	ctx.Status(http.StatusOK)
}
