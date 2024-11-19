package main

import (
	"final_api/routes"
	_ "fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server := gin.Default()
	routes.Router(server)
	server.Run(os.Getenv("PORT"))
}
