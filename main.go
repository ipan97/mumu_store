package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ipan97/mumu-store/app/models"
	"github.com/ipan97/mumu-store/config"
	"github.com/ipan97/mumu-store/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConfig := config.New(config.NewPostgreSQLConfig(), true)
	db, _ := dbConfig.DB()

	db.AutoMigrate(
		models.User{},
		models.Category{},
		models.Brand{},
		models.Product{},
	)

	s := server.Setup(db)
	httpServer := &http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
		Handler: s,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Print("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer db.Close()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("Timeout of 5 seconds.")
	default:
	}
	log.Println("Server exiting")
}
