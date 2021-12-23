package main

import (
	"fmt"
	"gin1gorm/models"
	"gin1gorm/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SSL"))
	// dsn := "host=192.168.1.2 port=5656 user=abhinav dbname=gorm1 password=Ab9450455997 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	if err != nil {
		println(err.Error())
	}
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("/register", routes.CreateUser(db))
	r.Run()
}
