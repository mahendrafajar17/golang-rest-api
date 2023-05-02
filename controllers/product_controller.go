package controllers

import (
	"database/sql"
	"net/http"

	"example.com/restapi/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	mysqlConnection := c.MustGet("mysqlConnection").(*sql.DB)
	products := models.GetProducts(mysqlConnection)

	if products == nil || len(products) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, products)
	}
}

func GetProduct(c *gin.Context) {
	mysqlConnection := c.MustGet("mysqlConnection").(*sql.DB)
	code := c.Param("code")

	product := models.GetProduct(mysqlConnection, code)

	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, product)
	}
}

func AddProduct(c *gin.Context) {
	mysqlConnection := c.MustGet("mysqlConnection").(*sql.DB)
	var product models.Product

	if err := c.BindJSON(&product); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddProduct(mysqlConnection, product)
		c.IndentedJSON(http.StatusCreated, product)
	}
}
