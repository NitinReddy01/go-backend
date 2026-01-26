package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NitinReddy01/go-backend/internal/config"
	"github.com/NitinReddy01/go-backend/internal/handler"
	"github.com/NitinReddy01/go-backend/internal/repository"
	"github.com/NitinReddy01/go-backend/internal/router"
	"github.com/NitinReddy01/go-backend/internal/server"
	"github.com/NitinReddy01/go-backend/internal/service"
)

func main() {
	config := config.LoadConfig()

	svr, err := server.New(config)

	if err != nil {
		log.Fatal("failed to initialise server:", err)
	}

	repos := repository.NewRepositories(svr.Pool)
	services := service.NewServices(repos)
	handlers := handler.NewHandlers(services, svr.Pool)

	router := router.New(handlers, config.CORSAllowedOrigins)

	svr.SetUpHTTPServer(router)

	signalCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := svr.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server", err)
		}
	}()

	<-signalCtx.Done()

	log.Println("Server shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := svr.Shutdown(shutdownCtx); err != nil {
		log.Print("server forced to shutdown: ", err)
	}

	log.Print("server exited properly")

}
