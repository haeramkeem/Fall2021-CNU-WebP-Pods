package controller

import (
    "pods/modules/logging"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

/**
 * Start Server
 * @params rootDir string: directory path for project root
 */
func Serve(rootDir string) {
	router := gin.Default()
    staticDir := rootDir + "/ui/dist"

    router.Use(static.Serve("/", static.LocalFile(staticDir, true)))
    router.NoRoute(func(c *gin.Context) {
        c.File(staticDir + "/index.html")
    })

    api := router.Group("/api")
    registerRouterGroup(api)

    logging.Log("Start PODs")
    if err := router.Run(); err != nil {
        panic(err)
    }
}

/**
 * Register handlers for API path
 * @params api *gin.RouterGroup: router to register
 */
func registerRouterGroup(api *gin.RouterGroup) {
    api.GET("/main/problem", getMainProblem)
    api.GET("/difficulty/problem", getDiffProblem)
    api.GET("/topic/problem", getTopicProblem)
    api.GET("/discussion/:probpath", getDiscussion)
}
