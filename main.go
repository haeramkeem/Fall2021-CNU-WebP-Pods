package main

import (
	"pods/controller"
	"pods/modules/logging"
)

func main() {

    // Middlewares
    logging.Init()

    // Start server
    controller.Serve()
}
