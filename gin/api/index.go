package api

import (
	"gin/database"
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAllIndex(c *gin.Context) {
	var products []models.Product

	database.GetDB().Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"list": products,
	})
}

type AddOneIndexRequest struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func AddOneIndex(c *gin.Context) {
	var request AddOneIndexRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product := models.Product{
		Name:  request.Name,
		Price: request.Price,
	}

	database.GetDB().Create(&product)

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}
