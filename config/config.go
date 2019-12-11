package config

import "time"

const (
	SentryTimeout time.Duration = 10

	DataBaseDriverName string = "mysql"

	DefaultDatabaseCharacterSet string = "utf8"

	TmpDirectory string = "/tmp/khazen"

	ImportMaxFile        int64  = 100 * 1024 * 1024 // 100MB
	ImportFileKey        string = "import_file"
	ImportTmpFilePattern string = "import-*.sql"
)

var StartTime time.Time

var Config = struct {
	Port    string `default:"3000"`
	LogFile string `default:"/var/log/khazen.log"`

	SentryDSN string

	AccessKey string `required:"true"`
	SecretKey string `required:"true"`

	MySQLCmd string `default:"mysql"`

	MySQL struct {
		Host     string `default:"127.0.0.1"`
		User     string `default:"root"`
		Password string `required:"true"`
		Port     string `default:"3306"`
	}
}{}
