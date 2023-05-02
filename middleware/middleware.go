package middleware

import (
	"database/sql"

	"example.com/restapi/controllers"
	"github.com/gin-gonic/gin"
)

func ApiMiddleware(mysqlConnection *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("mysqlConnection", mysqlConnection)
		c.Next()
	}
}

func SetUpRouter(mysqlConnection *sql.DB) *gin.Engine {
	router := gin.New()
	router.Use(ApiMiddleware(mysqlConnection))
	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:code", controllers.GetProduct)
	router.POST("/products", controllers.AddProduct)

	return router
}
