package main

import (
	"flag"
	"fmt"
	"strconv"

	"example.com/restapi/configurations"
	"example.com/restapi/middleware"
	"example.com/restapi/service"
)

func main() {
	yamlPath := flag.String("config", "config.yaml", "path config yaml file")
	config := configurations.LoadConfig(*yamlPath)
	fmt.Println(config)

	mysqlConnection := configurations.InitMySQL(config)
	router := middleware.SetUpRouter(mysqlConnection)

	// done := make(chan int, 1)
	Service := service.NewStompConService(config.Artemis.Host + ":" + strconv.Itoa(config.Artemis.Port))
	go func() {

		Service.Thread(5, "test")

		Service.Thread(10, "test")

	}()

	router.Run(config.App.Host + ":" + strconv.Itoa(config.App.Port))

	// <-done
	fmt.Println("Trying Stopped...")
}
