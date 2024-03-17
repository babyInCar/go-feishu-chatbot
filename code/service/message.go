package service

import (
	"context"
	. "feishu-chatbot/models"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/spf13/viper"
	"strings"
)

var client *lark.Client

type ReplyInfo struct {
	Title   string
	Content [][]map[string]interface{}
}

// GenMsg 生成向用户发送的消息
func GenMsg(str string, chatId string, messageId string, createTime string) (string, string, error) {

	// 把字母统一转成小写
	expression := strings.ToLower(str)
	//c := make(chan *string)
	var response string
	msgType := "text"
	//for {
	//	select {
	//	case receivedString := <-c:
	if strings.Contains(expression, "hello") {
		response = "Hi, What can I do for you?[Smile]"
	} else if strings.Contains(expression, "thank") {
		response = "It's my pleasure [Smile]"
	} else if strings.Contains(expression, "success") && strings.Contains(expression, "pay") {
		msgType = "post"
		response = "Congratulations, Pay success[CheckMark]!"
	} else if strings.Contains(expression, "fail") && strings.Contains(expression, "pay") {
		response = "Sorry, Pay Failed[CrossMark]，Please try again!"
	} else {
		response = "Sorry, I don't understand what you say![What?]"
	}
	err := CreateMessage(expression, chatId, messageId, createTime)
	if err != nil {
		fmt.Println(err)
		return "error", msgType, nil
	}
	return response, msgType, nil
	//case <-time.After(3 * time.Second):
	//	time.Sleep(time.Millisecond * 100)
	//default:

	//}
}

// CreateMessage 把用户输入的信息保存到数据库
func CreateMessage(content string, chatId string, messageId string, createTime string) error {
	message := Reply{Content: content, ChatId: chatId, MessageId: messageId, CreateTime: createTime}
	if err := message.Insert(); err != nil {
		// 插入数据表
		return err
	}
	return nil
}

// TransactionSucceed 接受外部系统的支付成功提示
func TransactionSucceed(message Reply) (string, string, error) {
	return GenMsg(message.Content, message.ChatId, message.MessageId, message.CreateTime)
}

func SendMsg(msg string, chatId *string, msgType string) {
	client = lark.NewClient(viper.GetString("APP_ID"),
		viper.GetString("APP_SECRET"))
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
