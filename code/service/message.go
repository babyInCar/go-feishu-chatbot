package service

import (
	. "feishu-chatbot/models"
)

func CreateMessage(content string, chatId string) error {
	message := Reply{Content: content, ChatId: chatId}
	if err := message.Insert(); err != nil {
		// 插入数据表
		return err
	}
	return nil
}
