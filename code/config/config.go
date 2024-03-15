package config

import (
	"encoding/json"
	"os"
)

// Config 整个项目的配置
type Config struct {
	Mode            string `json:"mode"`
	Host            string `json:"host"`
	Port            int    `json:"port"`
	SecretKey       string `json:"secret_key"`
	*LogConfig      `json:"log"`
	*DatabaseConfig `json:"database"`
}

type EmailConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	From string `json:"from"`
	User string `json:"username"`
	Pwd  string `json:"password"`
}

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver          string `json:"driver"`
	Host            string `json:"host"`
	Port            string `json:"port"`
	Database        string `json:"database"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Charset         string `json:"charset"`
	MaximumConn     int    `json:"maximum_connection"`
	MaximumFreeConn int    `json:"maximum_free_connection"`
	TimeOut         int    `json:"timeout"`
}

// Conf 全局配置变量
var Conf = &Config{}

func InitConfig(path string) error {
	/**
	filePath 配置文件json文件的路径
	*/
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, Conf)
	return err
}
