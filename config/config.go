package config

import "time"

const (
	SentryTimeout time.Duration = 10

	DefaultDatabaseCharacterSet = "utf8"

	TmpDirectory = "/tmp/khazen"

	ImportMaxFile        int64 = 100 * 1024 * 1024 // 100MB
	ImportFileKey              = "import_file"
	ImportTmpFilePattern       = "import-*.sql"

	SakkuUploadFileEndpoint = "/file/user/"
	SakkuUploadFileKeyFile  = "file"

	QueryTypeSelect    = 1
	QueryTypeNonSelect = 2
)

var StartTime time.Time

var ForbidenMySQLDatabaseNames = []string {
	"information_schema",
	"mysql",
	"performance_schema",
	"test",
	// sakku special
	"sakku",
	"sakku_core_db",
	"sakku_core_db_dev",
}
var ForbidenPostgresDatabaseNames = []string {
	"postgres", 
	"template2",
	"template1",
	"template0",
	"test",
	// sakku special
	"sakku",
}
var ForbidenMySQLUserNames = []string {
	"root", 
	"mysql",
	"admin",
	"test",
	// sakku special
	"sakku",
}
var ForbidenPostgresUserNames = []string {
	"postgres", 
	"root",
	"admin",
	"test",
	// sakku special
	"sakku",
}

var Config = struct {
	Port    string `default:"3000"`
	LogFile string `default:"/var/log/khazen.log"`

	SentryDSN string

	AccessKey string `required:"true"`
	SecretKey string `required:"true"`

	ServerZone string `default:"sakku.cloud"`

	UseSakkuUploadFileService bool `default:"false"`

	MySQLCmd     string `default:"mysql"`
	MySQLDumpCmd string `default:"mysqldump"`

	PostgresCmd     string `default:"psql"`
	PostgresDumpCmd string `default:"pg_dump"`

	MySQL struct {
		Host     string `default:"127.0.0.1"`
		User     string `default:"root"`
		Password string `required:"true"`
		Port     string `default:"3306"`
	}

	Postgres struct {
		Host     string `default:"127.0.0.1"`
		User     string `default:"postgres"`
		Password string `required:"true"`
		Port     string `default:"5432"`
	}

	SakkuUploadFile struct {
		Service    string `required:"true"`
		ServiceKey string `required:"true"`
	}
}{}
