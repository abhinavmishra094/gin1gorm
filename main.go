package main

import (
	"gin1gorm/models"
	"gin1gorm/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=192.168.1.2 port=5656 user=abhinav dbname=gorm1 password=Ab9450455997 sslmode=disable"
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
