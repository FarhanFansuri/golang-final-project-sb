package main

import (
	"final_api/models"
	"final_api/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/xo/dburl"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDB() {
	dsn, err := dburl.Parse("mysql://root:mOyIqTnSyxfccpttFWWWUeCFWKLgzsiw@autorack.proxy.rlwy.net:49673/railway") // sesuaikan dengan database kamu
	if err != nil {
		fmt.Println(err)
		return
	}
	DB, err = gorm.Open(mysql.Open(dsn.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database")
	}
	fmt.Println("Database connected successfully")
}

func main() {
	InitDB()
	// Migrate models ke database
	DB.AutoMigrate(&models.User{}, &models.Transaction{})

	err2 := godotenv.Load()
	if err2 != nil {
		log.Fatal("Error loading .env file")
	}
	server := gin.Default()
	routes.Router(server)
	server.Run(":8080")
}
