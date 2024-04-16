package app

import (
	"backend/internal/config"
	"log"
	"net/http"
)

func Run() error {

	if err := config.InitConfig(); err != nil {
		log.Panic(err)
	}

	serverUrl := config.CONFIG.Server.URL
	if err := http.ListenAndServe(serverUrl, nil); err != nil {
		log.Fatalf("server did not start work: %s", err.Error())
		return err
	}
	log.Println("Server listening url " + serverUrl)
	return nil
}
