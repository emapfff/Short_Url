package domain

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"test_ozon/config"
)

func Connect(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort)
	dbCon, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connect is fatal")
		return nil, err
	}

	log.Println("DB connected successfully!")

	return dbCon, nil
}

func Disconnect(db *gorm.DB) {
	dbCon, _ := db.DB()
	err := dbCon.Close()
	if err != nil {
		panic(err)
	}
}
