package main

import (
	"fmt"

	"example.com/restapi/configurations"
	"example.com/restapi/middleware"
	"example.com/restapi/service"
)

const host = "127.0.0.1"
const port = "61616"

func main() {
	mysqlConnection := configurations.InitMySQL()
	router := middleware.SetUpRouter(mysqlConnection)

	// done := make(chan int, 1)
	Service := service.NewStompConService(host + ":" + port)
	go func() {

		Service.Thread(5, "test")

		Service.Thread(10, "test")

	}()

	router.Run("localhost:8080")

	// <-done
	fmt.Println("Trying Stopped...")
}
