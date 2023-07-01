//go:build wireinject
// +build wireinject

package di

import (
	"avatg/internal/app/config"
	"avatg/internal/app/http"
	"avatg/internal/app/service"
	"avatg/internal/app/service/telegram"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	nethttp "net/http"
)

func InitializeApp(cfg config.Config) (*App, func(), error) {
	panic(wire.Build(
		AppSet,
	))
}

var AppSet = wire.NewSet(
	appConfigProvider,
	AppProvider,

	serverSet,
	TelegramSet,
)

var serverSet = wire.NewSet(
	ServerConfigProvider,
	http.ProvideServer,
	http.ProvideRouter,
	http.ProvidePingHandler,
	http.ProvideTelegramHandler,

	wire.Bind(new(nethttp.Handler), new(*gin.Engine)),
)

var TelegramSet = wire.NewSet(
	TelegramConfigProvider,
	telegram.ProvideTelegram,
	wire.Bind(new(service.Telegram), new(*telegram.Telegram)),
)

func appConfigProvider(cfg config.Config) AppConfig {
	return AppConfig{
		serverAddr: cfg.ServerAddr,
	}
}

func ServerConfigProvider(cfg config.Config) http.ServerConfig {
	return http.ServerConfig{
		Addr: cfg.ServerAddr,
	}
}

func TelegramConfigProvider(cfg config.Config) telegram.Config {
	return telegram.Config{}
}
