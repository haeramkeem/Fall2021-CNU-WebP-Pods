package controller

import (
	"github.com/gin-gonic/gin"
)

func Serve() error {
	router := gin.New()

    api := router.Group("/api")
    registerRouterGroup(api)

	router.Run()
    return nil
}

func registerRouterGroup(api *gin.RouterGroup) {
    api.GET("/main/problem", getMainProblem)
    api.GET("/main/discussion", getMainDiscuss)
    api.GET("/difficulty/problem", getDiffProblem)
    api.GET("/difficulty/discussion", getDiffDiscuss)
    api.GET("/topic/problem", getTopicProblem)
    api.GET("/topic/discussion", getTopicDiscuss)
}
