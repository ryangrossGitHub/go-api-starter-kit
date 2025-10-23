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
func GetItems(router *gin.Engine) {
	router.GET("/items", item_service.GetItems)
}

// CreateUser godoc
// @Summary Create a new item
// @Description Creates a new item with the provided details
// @Tags items
// @Accept json
// @Produce json
// @Param item body item_model.Item true "Item details for creation"
// @Success 201 {object} item_model.Item
// @Router /items [post]
func PostItem(router *gin.Engine) {
	router.POST("/items", item_service.PostItems)
}

