package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"strconv"
	"test_ozon/config"
	"test_ozon/domain"
	"test_ozon/domain/repository"
	"test_ozon/routes"
)

var Repo domain.UrlModel

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Loading .env is failed...")
	}
	cfg := config.Init()

	if cfg.StorageMethod == config.DB {
		db, _ := repository.Connect(&cfg.Database)
		repository.CreateUrlTable(db)
		defer repository.Disconnect(db)
	} else {
		repository.CreateUrlMap()
	}

	router := routes.SetupRoutes()

	port, err := strconv.Atoi(cfg.Server.Port)
	if err != nil {
		log.Fatal("Convert is fatal !")
	}
	log.Println("Server run!")
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), LogRequestMiddleware(router))
	if err != nil {
		log.Fatal(err)
	}
}
