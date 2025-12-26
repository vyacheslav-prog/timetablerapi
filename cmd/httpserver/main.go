package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"timetablerapi/internal/httpserver"
	"timetablerapi/internal/services"
)

const (
	srvAddr              = ":8080"
	srvReadHeaderTimeout = 15 * time.Second
	srvReadTimeout       = 10 * time.Second
	srvWriteTimeout      = 10 * time.Second
	srvIdleTimeout       = 30 * time.Second
)

func main() {
	ctx := context.Background()
	srvs, initErr := services.NewServices(ctx)
	if initErr != nil {
		log.Println("failed initialization for services:", initErr)
		os.Exit(1)
	}
	httpserver.NewMux(srvs)
	httpSrv := &http.Server{
		Addr:              srvAddr,
		ReadHeaderTimeout: srvReadHeaderTimeout,
		ReadTimeout:       srvReadTimeout,
		WriteTimeout:      srvWriteTimeout,
		IdleTimeout:       srvIdleTimeout,
	}
	if serveErr := httpSrv.ListenAndServe(); serveErr != nil {
		log.Println("failed serve for http-server:", serveErr)
	}
	log.Println("timetablerapi server for http listen 8080 port")
}
