package main

import (
	"log"
    "pods/service"
    _ "pods/controller"
)

func main() {
    logger := log.Default()
	logger.Println("App started")

    service.Fetch()

    //controller.Routing()
}
