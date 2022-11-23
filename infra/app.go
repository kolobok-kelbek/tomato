package infra

import (
	"github.com/kolobok-kelbek/go-example-service/infra/config"
	"github.com/kolobok-kelbek/go-example-service/infra/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	config config.Config
	logger logger.Logger
}

func NewApp(config config.Config, logger logger.Logger) *App {
	return &App{
		config: config,
		logger: logger,
	}
}

func (app *App) Run() error {
	app.logger.Infof("Start application")
	defer app.logger.Infof("Application stopped")

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	//// TODO: routing
	http.HandleFunc("/hello", app.getHello)

	go func() {
		err := http.ListenAndServe(app.config.Server.Port, nil)
		if err != nil {
			panic(err)
		}
	}()

	<-exit

	return nil
}

func (app *App) getHello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello world"))
	if err != nil {
		return
	}
}
