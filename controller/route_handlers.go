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
    date := ctx.DefaultQuery("date", "default date")
    problem := ProblemMock(date)

    sendResponse(ctx, nil, problem)
}

func getMainDiscuss(ctx *gin.Context) {
    date := ctx.DefaultQuery("date", "default date")
    lang := ctx.DefaultQuery("lang", "default language")
    discuss := DiscussionMock(date, lang)

    sendResponse(ctx, nil, discuss)
}

func getDiffProblem(ctx *gin.Context) {
    level := ctx.DefaultQuery("level", "default level")
    problem := ProblemMock(level)

    sendResponse(ctx, nil, problem)
}

func getDiffDiscuss(ctx *gin.Context) {
    level := ctx.DefaultQuery("level", "default level")
    lang := ctx.DefaultQuery("lang", "default language")
    discuss := DiscussionMock(level, lang)

    sendResponse(ctx, nil, discuss)
}

func getTopicProblem(ctx *gin.Context) {
    topic := ctx.DefaultQuery("topic", "default topic")
    problem := ProblemMock(topic)

    sendResponse(ctx, nil, problem)
}

func getTopicDiscuss(ctx *gin.Context) {
    topic := ctx.DefaultQuery("topic", "default topic")
    lang := ctx.DefaultQuery("lang", "default language")
    discuss := DiscussionMock(topic, lang)

    sendResponse(ctx, nil, discuss)
}
