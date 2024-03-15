package models

import . "feishu-chatbot/database"

type Reply struct {
	Content string `json:"content"`
	//MsgId     string `json:"msg_id"`
	ChatId    string `json:"chat_id" comment:"用户的Id"`
	ReplyTime string `json:"reply_time" comment:"回复时间"`
}

func (r Reply) TableName() string {
	return "reply"
}

func (r *Reply) Insert() error {
	err := Orm.Create(&r).Error
	return err
}
