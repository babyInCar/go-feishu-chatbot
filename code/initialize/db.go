package initialize

import (
	"feishu-chatbot/config"
	"feishu-chatbot/database"
	"feishu-chatbot/models"
)

func InitDB(cfg *config.DatabaseConfig) {
	Orm := database.GetOrm(cfg)
	// 禁用复数
	Orm.SingularTable(true)
	Orm.AutoMigrate(&models.Reply{})
}
