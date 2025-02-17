package main

import (
	item_controller "test-api/controller"

	_ "test-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Item API
func main() {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	item_controller.Setup(router)

	router.Run("localhost:8080")
}
