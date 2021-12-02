package controller

import (
	. "pods/service"

	"github.com/gin-gonic/gin"
)

func sendResponse(ctx *gin.Context, err error, content interface{}) {
    if err != nil {
        ctx.JSON(200, gin.H{"error": err.Error(), "content": content})
    } else {
        ctx.JSON(200, gin.H{"error": nil, "content": content})
    }
}

func getMainProblem(ctx *gin.Context) {
    date := ctx.Query("date") 
    problem := GetMainProb(date)

    sendResponse(ctx, nil, problem)
}

func getDiffProblem(ctx *gin.Context) {
    level := ctx.Query("level")
    problem := GetDifficultyProb(level)

    sendResponse(ctx, nil, problem)
}

func getTopicProblem(ctx *gin.Context) {
    topic := ctx.Query("topic")
    problem := GetTopicProb(topic)

    sendResponse(ctx, nil, problem)
}


func getDiscussion(ctx *gin.Context) {
    problemPath := ctx.Param("probpath")
    lang := ctx.Query("lang")

    discussion, err := GetDiscussion(problemPath, lang)

    sendResponse(ctx, err, discussion)
}
