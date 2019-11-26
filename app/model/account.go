package model

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"khazen/config"
)

type MySQLAccount struct {
	Username              string `json:"username"`
	Password              string `json:"password"`
	MaxQueriesPerHour     string `json:"max_queries_per_hour"`
	MaxUpdatesPerHour     string `json:"max_updates_per_hour"`
	MaxConnectionsPerHour string `json:"max_connections_per_hour"`
	MaxUserConnections    string `json:"max_user_connections"`
	NativePassword        bool   `json:"native_password"`
}

func (account *MySQLAccount) Create(dbURI string) (err error) {
	if account.hasRequirements() {
		queryUsername := "CREATE USER IF NOT EXISTS ?@'%'"
		queryPassword := queryUsername + " IDENTIFIED BY ?"
		if account.NativePassword {
			queryPassword = queryUsername + " IDENTIFIED WITH mysql_native_password BY ?"
		}
		query := queryPassword + " WITH MAX_QUERIES_PER_HOUR ? MAX_UPDATES_PER_HOUR ? MAX_CONNECTIONS_PER_HOUR ? MAX_USER_CONNECTIONS ?"
		db, err := sql.Open("mysql", dbURI)
		if err == nil {
			defer db.Close()
			_, err = db.Exec(query, account.Username, account.Password, account.MaxQueriesPerHour, account.MaxUpdatesPerHour, account.MaxConnectionsPerHour, account.MaxUserConnections)
		}
	} else {
		err = errors.New(config.InvalidArgsMessage)
	}
	return
}

func (account *MySQLAccount) hasRequirements() bool {
	if account.Username == "" || account.Password == "" || account.MaxQueriesPerHour == "" || account.MaxUpdatesPerHour == "" || account.MaxConnectionsPerHour == "" || account.MaxUserConnections == "" {
		return false
	}
	return true
}
