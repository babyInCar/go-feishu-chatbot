package main

import (
	"context"
	"encoding/json"
	"fmt"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/core/httpserverext"
	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	"log"
	"net/http"
	"regexp"

	"github.com/spf13/viper"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"

	lark "github.com/larksuite/oapi-sdk-go/v3"
)

func setEnv() {
	viper.SetConfigFile("./feishu_config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

var client *lark.Client

func init() {
	setEnv()
}

func sendMsg(msg string, chatId *string) {
	content := larkim.NewTextMsgBuilder().
		Text(msg).
		Build()

	resp, err := client.Im.Message.Create(context.Background(), larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(larkim.ReceiveIdTypeChatId).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			MsgType(larkim.MsgTypeText).
			ReceiveId(*chatId).
			Content(content).
			Build()).
		Build())

	// 处理错误
	if err != nil {
		fmt.Println(err)
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
	}
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
	log.Printf("content is:%s", content)
	err := json.Unmarshal([]byte(content), &contentMap)
	if err != nil {
		fmt.Println(err)
	}
	text := contentMap["text"].(string)
	return msgFilter(text)
}

//func main() {
//	client = lark.NewClient(viper.GetString("APP_ID"),
//		viper.GetString("APP_SECRET"))
//
//	//// 注册消息处理器
//	handler := dispatcher.NewEventDispatcher(viper.GetString(
//		"APP_VERIFICATION_TOKEN"), viper.GetString("APP_ENCRYPT_KEY")).
//		OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
//			// 处理消息 event，这里简单打印消息的内容
//			fmt.Println(larkcore.Prettify(event))
//			fmt.Println(event.RequestId())
//			return nil
//		}).OnP2MessageReadV1(func(ctx context.Context, event *larkim.P2MessageReadV1) error {
//		// 处理消息 event，这里简单打印消息的内容
//		fmt.Println(larkcore.Prettify(event))
//		fmt.Println(event.RequestId())
//		return nil
//	})
//
//	r := gin.Default()
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	})
//
//	// 在已有 Gin 实例上注册消息处理路由
//	r.POST("/webhook/event", sdkginext.NewEventHandlerFunc(handler))
//
//	fmt.Println("http server started",
//		"http://localhost:9000/webhook/event")
//
//	r.Run(":9000")
//
//}

func main() {
	// 注册消息处理器
	handler := dispatcher.NewEventDispatcher("verificationToken", "eventEncryptKey").OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
		// 处理消息 event，这里简单打印消息的内容
		fmt.Println(larkcore.Prettify(event))
		fmt.Println(event.RequestId())
		return nil
	}).OnP2MessageReadV1(func(ctx context.Context, event *larkim.P2MessageReadV1) error {
		// 处理消息 event，这里简单打印消息的内容
		fmt.Println(larkcore.Prettify(event))
		fmt.Println(event.RequestId())
		return nil
	})

	// 注册 http 路由
	http.HandleFunc("/webhook/event", httpserverext.NewEventHandlerFunc(handler, larkevent.WithLogLevel(larkcore.LogLevelDebug)))

	// 启动 http 服务
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}
