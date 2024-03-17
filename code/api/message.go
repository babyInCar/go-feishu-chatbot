package api

import (
	. "feishu-chatbot/models"
	"feishu-chatbot/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetPayResult 获取外部系统传过来的支付情况
func GetPayResult(ctx *gin.Context) {
	message := Reply{}
	//chatId := ctx.ShouldBindJSON(&service.Reply{})
	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    1000,
			"message": "server Internal error",
		})
	}
	response, msgType, err := service.TransactionSucceed(message)

	err = service.SendMsg(ctx, response, &message.ChatId, msgType)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "send message result fail!",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "send message result Success!",
	})
}
