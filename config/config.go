package config

import "time"

const (
	SentryTimeout time.Duration = 10

	InvalidArgsMessage string = "invalid args"
)

var Config = struct {
	Port    string `default:"3000"`
	LogFile string `default:"/var/log/khazen.log"`

	SentryDSN string

	AccessKey string `required:"true" env:"AccessKey"`
	SecretKey string `required:"true" env:"SecretKey"`

	MySQL struct {
		Host     string `default:"127.0.0.1" env:"MysqlHost"`
		User     string `default:"root" env:"MysqlUser"`
		Password string `required:"true" env:"MysqlPassword"`
		Port     string `default:"3306" env:"MysqlPort"`
	}
}{}
