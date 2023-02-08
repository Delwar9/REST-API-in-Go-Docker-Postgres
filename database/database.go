package database

import (
	"fmt"
	"log"
	"os"

	"github.com/delwar/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect database.", err)
		os.Exit(2)
	}
	log.Println("database connection successful.")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Printf("running migrations...")
	db.AutoMigrate(&model.Fact{})

	DB = Dbinstance{Db: db}
}
