package infra

import (
	"github.com/kolobok-kelbek/go-example-service/infra/config"
	"net/http"
)

type App struct {
	config *config.Config
}

func NewApp(config *config.Config) *App {
	return &App{
		config: config,
	}
}

func (app *App) Run() {
	// TODO: routing
	http.HandleFunc("/hello", app.getHello)

	err := http.ListenAndServe(app.config.Server.Port, nil)
	if err != nil {
		panic(err)
	}
}

func (app *App) getHello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello world"))
	if err != nil {
		return
	}
}
