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
    api.GET("/difficulty/problem", getDiffProblem)
    api.GET("/topic/problem", getTopicProblem)
    api.GET("/discussion/:probpath", getDiscussion)
}
