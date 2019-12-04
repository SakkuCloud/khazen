package model

type MySQLDatabase struct {
	Username string `json:"username"`
	Database string `json:"database"`
}

func (database *MySQLDatabase) HasRequirements() bool {
	if database.Username == "" || database.Database == "" {
		return false
	}
	return true
}

func (database *MySQLDatabase) GetCreateQuery() (query string) {
	query = "CREATE DATABASE IF NOT EXISTS " + database.Database
	return
}

func (database *MySQLDatabase) GetSetPrivilegesQuery() (query string) {
	query = "GRANT ALL PRIVILEGES ON " + database.Database + ".* TO '" + database.Username + "'@'%'"
	return
}
