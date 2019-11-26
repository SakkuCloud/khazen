package model

import (
	"database/sql"
	"errors"
	"khazen/config"
)

type MySQLDatabase struct {
	Username string `json:"username"`
	Database string `json:"database"`
}

func (database *MySQLDatabase) Create(dbURI string) (err error) {
	if database.hasRequirements() {
		query := "CREATE DATABASE IF NOT EXISTS ?"
		db, err := sql.Open("mysql", dbURI)
		if err == nil {
			defer db.Close()
			_, err = db.Exec(query, database.Database)
		}
	} else {
		err = errors.New(config.InvalidArgsMessage)
	}
	return
}

func (database *MySQLDatabase) SetPrivileges(dbURI string) (err error) {
	if database.hasRequirements() {
		query := "GRANT ALL PRIVILEGES ON ?.* TO ?@'%'"
		db, err := sql.Open("mysql", dbURI)
		if err == nil {
			defer db.Close()
			_, err = db.Exec(query, database.Database, database.Username)
		}
	} else {
		err = errors.New(config.InvalidArgsMessage)
	}
	return
}

func (database *MySQLDatabase) hasRequirements() bool {
	if database.Username == "" || database.Database == "" {
		return false
	}
	return true
}
