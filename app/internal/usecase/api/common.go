package api

import (
	"github.com/igorshmel/lic_auto_post/app/pkg/errs"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
)

func extErr(code string, msg string, log logger.Logger) error {
	log.Error(msg)
	return errs.New().SetCode(code).SetMsg(msg).Error()
}
