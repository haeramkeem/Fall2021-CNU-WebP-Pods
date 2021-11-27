package main

import (
	"log"
    "pods/service"
    "pods/controller"
)

func main() {
    logger := log.Default()
	logger.Println("App started")

    service.Fetch()

    controller.Routing()
}
