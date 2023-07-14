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

// CreateRecordEndpoint --
type CreateRecordEndpoint struct {
	log     logger.Logger
	usecase port.CreateRecordUseCase
}

// NewCreateRecordEndpoint _
func NewCreateRecordEndpoint(usecase port.CreateRecordUseCase, log logger.Logger) port.CreateRecordEndpoint {
	return CreateRecordEndpoint{
		log:     log,
		usecase: usecase,
	}
}

// CreateRecordExecute is handler
func (ths CreateRecordEndpoint) CreateRecordExecute(ctx *gin.Context) {
	ths.log = middleware.SetRequestIDPrefix(ctx, ths.log)
	log := ths.log.WithMethod("endpoint create record")

	req := dto.CreateRecordReqDTO{}

	// request parse
	if err := req.Parse(ctx); err != nil {
		log.Error("unable to parse request: %s", err)
		ctx.JSON(http.StatusOK, errs.New().SetCode(errs.ParseRequest).
			SetMsg("unable to parse a request").
			GinJSON())
		return
	}

	log.Debug("req: %v", req)

	// validate request
	if err := req.Validate(); err != nil {
		log.Error("error of validation: %s", err)
		ctx.JSON(http.StatusOK, errs.New().SetCode(errs.Syntax).
			SetMsg(err.Error()).
			GinJSON())
		return
	}

	err := ths.usecase.Execute(ctx, &req)
	if err != nil {
		log.Error("failed call to usecase CreateRecord: %s", err)
		ctx.JSON(http.StatusOK, errs.FromError(err).GinJSON())
		return
	}

	ctx.Status(http.StatusOK)
}
