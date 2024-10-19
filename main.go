package main

import (
	"github.com/joho/godotenv"
	"log"
	"test_ozon/config"
	"test_ozon/domain"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Loading .env is failed...")
	}
	cfg := config.Init().Database
	db, _ := domain.Connect(&cfg)

	defer domain.Disconnect(db)
}
