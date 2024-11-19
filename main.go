package main

import (
	"final_api/models"
	"final_api/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local" // sesuaikan dengan database kamu
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
	server.Run(os.Getenv("PORT"))
}
