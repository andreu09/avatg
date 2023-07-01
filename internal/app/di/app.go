package di

import (
	"avatg/internal/app/http"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type AppConfig struct {
	serverAddr string
}

type App struct {
	cfg    AppConfig
	server *http.Server
	err    chan error
}

func (app *App) Run(cleanupResourcesFunc func()) {
	go app.RunInterruptHandler(cleanupResourcesFunc)
	go app.RunServer()

	log.Println("exit", <-app.err)
}

func (app *App) RunServer() {
	app.err <- app.server.Run()
}

func (app *App) RunInterruptHandler(cleanupResourcesFunc func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	terminateError := fmt.Errorf("%s", <-c)

	cleanupResourcesFunc()
	app.err <- terminateError
}

func AppProvider(
	cfg AppConfig,
	server *http.Server,
) (*App, func()) {
	return &App{
			cfg:    cfg,
			err:    make(chan error),
			server: server,
		}, func() {
			server.Stop()
		}
}
