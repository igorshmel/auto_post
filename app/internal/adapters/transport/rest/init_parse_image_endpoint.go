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

// InitParseImageEndpoint --
type InitParseImageEndpoint struct {
	log     logger.Logger
	usecase port.InitParseImageUseCase
}

// NewInitParseImageEndpoint _
func NewInitParseImageEndpoint(usecase port.InitParseImageUseCase, log logger.Logger) port.Endpoint {
	return InitParseImageEndpoint{
		log:     log,
		usecase: usecase,
	}
}

// Execute is handler
func (ths InitParseImageEndpoint) Execute(ctx *gin.Context) {
	ths.log = middleware.SetRequestIDPrefix(ctx, ths.log)
	log := ths.log.WithMethod("endpoint InitParseImage")

	req := dto.NewParseImageReq()

	// request parse
	if err := req.Parse(ctx); err != nil {
		log.Error("unable to parse a request: %s", err)
		ctx.JSON(http.StatusOK, errs.New().SetCode(errs.ParseRequest).
			SetMsg("unable to parse a request").
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
		log.Error("failed call to usecase InitParseImage: %s", err)
		ctx.JSON(http.StatusOK, errs.FromError(err).GinJSON())
		return
	}

	ctx.Status(http.StatusOK)
}
