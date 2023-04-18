package api

import (
	"auto_post/app/pkg/errs"
	logger "auto_post/app/pkg/log"
)

func extErr(code string, msg string, log logger.Logger) error {
	log.Error(msg)
	return errs.New().SetCode(code).SetMsg(msg).Error()
}
