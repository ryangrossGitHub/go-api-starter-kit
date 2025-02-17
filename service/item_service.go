package item_service

import (
	"net/http"
	item_model "test-api/model"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	items := []item_model.Item{
		{Id: 1, Title: "Halo", Price: 22.99, ReleaseDate: 1736501470, Rating: "T"},
		{Id: 2, Title: "Halo 2", Price: 35.99, ReleaseDate: 1736601470, Rating: "T"},
	}

	c.IndentedJSON(http.StatusOK, items)
}
