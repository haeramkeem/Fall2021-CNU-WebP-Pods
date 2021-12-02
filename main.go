package main

import (
	"os"
	"pods/controller"
)

func main() {
    // Configuration
    root, _ := os.Getwd()

    // Start server
    controller.Serve(root)
}
