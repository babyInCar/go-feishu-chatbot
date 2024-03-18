package main

import (
	"context"
	"encoding/json"
	"feishu-chatbot/api"
	"feishu-chatbot/initialize"
	"feishu-chatbot/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"path/filepath"
	"regexp"

	sdkginext "github.com/larksuite/oapi-sdk-gin"

	. "feishu-chatbot/config"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func setEnv() {
	viper.SetConfigFile("./feishu_config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

//var client *lark.Client

func init() {
	setEnv()
}

func msgFilter(msg string) string {
	//replace @到下一个非空的字段 为 ''
	regex := regexp.MustCompile(`@[^ ]*`)
	return regex.ReplaceAllString(msg, "")

}
func parseContent(content string) string {
	//"{\"text\":\"@_user_1  hahaha\"}",
	//only get text content hahaha
	var contentMap map[string]interface{}
	err := json.Unmarshal([]byte(content), &contentMap)
	if err != nil {
		fmt.Println(err)
	}
	text := contentMap["text"].(string)
	return msgFilter(text)
}
func main() {
	//client = lark.NewClient(viper.GetString("APP_ID"),
	//	viper.GetString("APP_SECRET"))
	// 注册消息处理器
	handler := dispatcher.NewEventDispatcher(viper.GetString(
		"APP_VERIFICATION_TOKEN"), viper.GetString("APP_ENCRYPT_KEY")).
		OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
			fmt.Println(larkcore.Prettify(event))
			content := event.Event.Message.Content
			contentStr := parseContent(*content)
			out, msgType, err := service.GenMsg(contentStr, *event.Event.Message.ChatId, *event.Event.Message.MessageId, *event.Event.Message.CreateTime)
			fmt.Printf("out is: %s\n", out)
			if err != nil {
				fmt.Println(err)
			}
			go func() {
				_ = service.SendMsg(ctx, out, event.Event.Message.ChatId, msgType)
			}()
			return nil
		})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 处理外部的交易连接请求，可用PostMan测试
	r.POST("/transaction", api.GetPayResult)
	// 初始化相关配置
	dir, err := filepath.Abs(filepath.Dir("./"))
	if err != nil {
		fmt.Println("error is:", err)
	}
	if err := InitConfig(fmt.Sprintf("%s/config.json", dir)); err != nil {
		fmt.Println("error is:", err)
		panic(err.Error())
	}
	// 初始化Mysql的连接
	initialize.InitDB(Conf.DatabaseConfig)

	// 在已有 Gin 实例上注册消息处理路由
	r.POST("/webhook/event", sdkginext.NewEventHandlerFunc(handler))

	fmt.Println("http server started",
		"http://localhost:9000/webhook/event")

	r.Run(":9000")

}
