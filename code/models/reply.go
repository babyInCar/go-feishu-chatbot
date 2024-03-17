package models

import (
	. "feishu-chatbot/database"
)

type Reply struct {
	Content    string `json:"content"`
	ChatId     string `json:"chat_id" comment:"用户的Id"`
	CreateTime string `json:"create_time" comment:"发送消息时间"`
	MessageId  string `json:"message_id" comment:"信息id"`
}

func (r Reply) TableName() string {
	return "reply"
}

func (r *Reply) Insert() error {
	err := Orm.Create(&r).Error
	return err
}
