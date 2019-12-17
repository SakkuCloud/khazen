package model

import (
	_ "github.com/go-sql-driver/mysql"
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

func (account *MySQLAccount) HasRequirements() bool {
	if account.Username == "" || account.Password == "" || account.MaxQueriesPerHour == "" || account.MaxUpdatesPerHour == "" || account.MaxConnectionsPerHour == "" || account.MaxUserConnections == "" {
		return false
	}
	return true
}

func (account *MySQLAccount) GetCreateQuery() (query string) {
	query = "CREATE USER IF NOT EXISTS '" + account.Username + "'@'%'"
	query = query + " IDENTIFIED BY '" + account.Password + "'"
	query = query + " WITH MAX_QUERIES_PER_HOUR " + account.MaxQueriesPerHour
	query = query + " MAX_UPDATES_PER_HOUR " + account.MaxUpdatesPerHour
	query = query + " MAX_CONNECTIONS_PER_HOUR " + account.MaxConnectionsPerHour
	query = query + " MAX_USER_CONNECTIONS " + account.MaxUserConnections
	return
}

type PostgresAccount struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConnectionLimit string `json:"connection_limit"`
}

func (account *PostgresAccount) HasRequirements() bool {
	if account.Username == "" || account.Password == "" || account.ConnectionLimit == "" {
		return false
	}
	return true
}

func (account *PostgresAccount) GetCreateQuery() (query string) {
	query = "CREATE ROLE " + account.Username
	query = query + " WITH ENCRYPTED PASSWORD '" + account.Password + "'"
	query = query + " LOGIN CONNECTION LIMIT " + account.ConnectionLimit
	return
}
