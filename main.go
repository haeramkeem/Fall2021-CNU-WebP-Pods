package main

import (
	"os"
	"pods/controller"
)

func main() {
    // Get directory path for root of project
    root, _ := os.Getwd()

    // Start server
    controller.Serve(root)
}
