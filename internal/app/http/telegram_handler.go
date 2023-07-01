package http

import (
	"avatg/internal/app/service"
	"github.com/gin-gonic/gin"
)

type TelegramHandler struct {
	TelegramSvc service.Telegram
}

func ProvideTelegramHandler(telegramSvc service.Telegram) *TelegramHandler {
	h := new(TelegramHandler)
	h.TelegramSvc = telegramSvc

	return h
}

func (h TelegramHandler) Start(ctx *gin.Context) {
	h.TelegramSvc.Run()
}
