package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"open-market.com/user-api/cmd/api/routes"
	"open-market.com/user-api/pkg/config"
	"open-market.com/user-api/pkg/helpers"
)

func main() {
	cfg, err := config.Load()
	helpers.CheckErr(err)

	r := gin.Default()

	// db := db.DB()
	// defer db.Close()

	routes.SetupRoutes(r)

	srv := &http.Server{
		Addr:    ":" + cfg.App.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
