package controller

import (
	. "pods/service"

	"github.com/gin-gonic/gin"
)

/**
 * Send response to client
 * @params ctx *gin.Context: serving context
 * @params err error: generated error if exists
 * @params content interface{}: content to reply
 */
func sendResponse(ctx *gin.Context, err error, content interface{}) {
    if err != nil {
        ctx.JSON(200, gin.H{"error": err.Error(), "content": content})
    } else {
        ctx.JSON(200, gin.H{"error": nil, "content": content})
    }
}

/**
 * Send main problem to client
 * @params ctx *gin.Context: serving context
 */
func getMainProblem(ctx *gin.Context) {
    date := ctx.Query("date") 
    problem, err := GetMainProb(date)

    sendResponse(ctx, err, problem)
}

/**
 * Send problem selected by requested difficulty to client
 * @params ctx *gin.Context: serving context
 */
func getDiffProblem(ctx *gin.Context) {
    level := ctx.Query("level")
    problem, err := GetDifficultyProb(level)

    sendResponse(ctx, err, problem)
}

/**
 * Send problem selected by requested algorithm topic to client
 * @params ctx *gin.Context: serving context
 */
func getTopicProblem(ctx *gin.Context) {
    topic := ctx.Query("topic")
    problem, err := GetTopicProb(topic)

    sendResponse(ctx, err, problem)
}

/**
 * Send discussions of problem to client
 * @params ctx *gin.Context: serving context
 */
func getDiscussion(ctx *gin.Context) {
    problemPath := ctx.Param("probpath")
    lang := ctx.Query("lang")

    discussion, err := GetDiscussion(problemPath, lang)

    sendResponse(ctx, err, discussion)
}
