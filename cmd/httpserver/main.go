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
	log.Println("timetablerapi services is initializing...")
	srvs, initErr := services.NewServices(ctx)
	if initErr != nil {
		log.Println("timetablerapi services is failed:", initErr)
		os.Exit(1)
	}
	log.Println("timetablerapi services is initialized")
	httpserver.NewMux(srvs)
	httpSrv := &http.Server{
		Addr:              srvAddr,
		ReadHeaderTimeout: srvReadHeaderTimeout,
		ReadTimeout:       srvReadTimeout,
		WriteTimeout:      srvWriteTimeout,
		IdleTimeout:       srvIdleTimeout,
	}
	log.Println("timetablerapi a http server is listening addr:", srvAddr)
	if serveErr := httpSrv.ListenAndServe(); serveErr != nil {
		log.Println("timetablerapi a http server is failed:", serveErr)
		os.Exit(1)
	}
}
