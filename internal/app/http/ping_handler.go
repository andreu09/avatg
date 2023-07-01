package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingHandler struct {
}

func ProvidePingHandler() *PingHandler {
	h := new(PingHandler)

	return h
}

func (h PingHandler) Handle(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
