package item_controller

import (
	item_service "test-api/service"

	"github.com/gin-gonic/gin"
)

// GetItems             godoc
// @Summary      Get all items
// @Description  Responds with the list of all items as JSON.
// @Tags         items
// @Produce      json
// @Success      200 {array} item_model.Item
// @Router       /items [get]
func Setup(router *gin.Engine) {
	router.GET("/items", item_service.GetItems)
}
