package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"strconv"
	"test_ozon/config"
	"test_ozon/domain"
	"test_ozon/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Loading .env is failed...")
	}
	cfg := config.Init()
	db, _ := domain.Connect(&cfg.Database)

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
	defer domain.Disconnect(db)
}
