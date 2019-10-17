package main

import (
	"github.com/FernandoCagale/c4-active/api/middleware"
	"github.com/FernandoCagale/c4-active/pkg/infra/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var log *zap.SugaredLogger

func init() {
	godotenv.Load()

	log = logger.Get()
}

func main() {
	defer log.Sync()

	log.Debug("Initializing actives")

	app, e := SetupApplication()

	if e != nil {
		panic("Erro to start application")
	}

	router := app.MakeHandlers()

	router.Use(middleware.Header)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Debug("Started successfully")

	log.Fatal(srv.ListenAndServe())
}
