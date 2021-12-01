package controller

import (
	"pods/modules/config"

	"github.com/gin-gonic/gin"
)

func Serve() error {
	router := gin.Default()

    staticDir := config.Get("staticdir")
    router.Static("/", staticDir)

    api := router.Group("/api/")
    registerRouterGroup(api)

	router.Run()
    return nil
}

func registerRouterGroup(api *gin.RouterGroup) {

}
