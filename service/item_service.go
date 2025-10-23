package item_service

import (
	"net/http"
	item_model "test-api/model"

	"github.com/gin-gonic/gin"
)

var items = []item_model.Item{
	{Id: 1, Title: "Halo", Price: 22.99, ReleaseDate: 1736501470, Rating: "T"},
	{Id: 2, Title: "Halo 2", Price: 35.99, ReleaseDate: 1736601470, Rating: "T"},
}

func GetItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func PostItems(c *gin.Context) {
	var item item_model.Item

	if err := c.BindJSON(&item); err != nil {
		return
	}

	items = append(items, item)
	c.IndentedJSON(http.StatusCreated, item)
}
