package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"report-service/handler"
	mw "report-service/middleware"
)

func main() {
	_ = godotenv.Load()
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(mw.AuthMiddleware)
	r.Use(mw.LoggingMiddleware)

	r.Get("/api/v1/students/{id}/report", handler.StudentReportHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Infof("Starting server on :%s", port)
	http.ListenAndServe(":"+port, r)
}
