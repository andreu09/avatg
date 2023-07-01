package http

import (
	"github.com/gin-gonic/gin"
)

func ProvideRouter(
	pingHandler *PingHandler,
	telegramHandler *TelegramHandler,
) *gin.Engine {
	router := gin.New()

	router.GET("ping", pingHandler.Handle)
	router.GET("telegram/start", telegramHandler.Start)

	return router
}
