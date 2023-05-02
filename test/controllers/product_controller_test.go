package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/restapi/configurations"
	"example.com/restapi/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func setup() {
	mysqlConnection := configurations.InitMySQL()
	router = middleware.SetUpRouter(mysqlConnection)
}

func TestHomepageHandler(t *testing.T) {
	setup()
	// mockResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`
	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// responseData, _ := ioutil.ReadAll(w.Body)
	// assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
