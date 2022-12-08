package api

import (
	"git.fintechru.org/dfa/dfa_lib/errs"
	"git.fintechru.org/dfa/dfa_lib/logger"
)

func extErr(code string, msg string, log logger.Logger) error {
	log.Error(msg)
	return errs.New().SetCode(code).SetMsg(msg).Error()
}
