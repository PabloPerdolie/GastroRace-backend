package app

import (
	"backend/internal/config"
	"backend/internal/handlers"
	"backend/internal/mongodb"
	"context"
	"log"
	"net/http"
	"time"
)

func Run() error {

	if err := config.InitConfig(); err != nil {
		log.Panic(err)
	}

	if err := mongodb.InitDB(context.Background()); err != nil {
		log.Panic(err)
	}

	go func() {
		for {
			time.Sleep(5 * time.Minute)
			log.Println("Server is running and ready to accept connections")
		}
	}()

	serverUrl := config.CONFIG.Server.URL
	if err := http.ListenAndServe(serverUrl, handlers.SetupRoutes()); err != nil {
		log.Fatalf("server did not start work: %s", err.Error())
		return err
	}
	log.Println("Server listening url " + serverUrl)
	return nil
}
