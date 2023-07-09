package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tniita/go-backend-mysql/crud"
	"github.com/tniita/go-backend-mysql/repo"
)

func GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, crud.GetItems())
}

func GetItemById(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusOK, crud.GetItemById(id))
}

func UpdateItem(c *gin.Context) {
	var item repo.Items
	if err := c.BindJSON(&item); err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, crud.UpdateItem(item))
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusOK, crud.DeleteItem(id))
}
