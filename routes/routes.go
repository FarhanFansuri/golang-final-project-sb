package routes

import (
	"final_api/controllers"

	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {
	server.GET("/", controllers.GetData)
}
