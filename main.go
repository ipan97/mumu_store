package main

import (
	"context"
	"github.com/ipan97/mumu-store/app/models"
	"github.com/ipan97/mumu-store/config"
	"github.com/ipan97/mumu-store/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	dbConfig := config.Config{
		Database:      config.NewPostgreSQLConfig(),
		IsDevelopment: true,
	}
	db, _ := dbConfig.GetDB()
	db.AutoMigrate(models.User{}, models.Category{}, models.Brand{}, models.Product{})

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
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Print("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
