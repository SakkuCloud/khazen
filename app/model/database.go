package model

import (
	"khazen/util"
	"khazen/config"
)

type MySQLDatabase struct {
	Username     string `json:"username"`
	Database     string `json:"database"`
	CharacterSet string `json:"character_set"`
}

func (database *MySQLDatabase) HasRequirements() bool {
	if database.Username == "" || database.Database == "" {
		return false
	}
	if util.ArrayContains(config.ForbidenMySQLDatabaseNames,database.Database) {
		return false
	}
	return true
}

func (database *MySQLDatabase) SetDefaults() {
	if database.CharacterSet == ""{
		database.CharacterSet = config.DefaultDatabaseCharacterSet
	}
}

func (database *MySQLDatabase) GetCreateQuery() (query string) {
	query = "CREATE DATABASE IF NOT EXISTS " + database.Database + " CHARACTER SET = '"+ database.CharacterSet +"'"
	return
}

func (database *MySQLDatabase) GetSetPrivilegesQuery() (query string) {
	query = "GRANT ALL PRIVILEGES ON " + database.Database + ".* TO '" + database.Username + "'@'%'"
	return
}

func (database *MySQLDatabase) GetDeleteQuery() (query string) {
	query = "DROP DATABASE IF EXISTS " + database.Database
	return
}

type PostgresDatabase struct {
	Username     string `json:"username"`
	Database     string `json:"database"`
}

func (database *PostgresDatabase) HasRequirements() bool {
	if database.Username == "" || database.Database == "" {
		return false
	}
	if util.ArrayContains(config.ForbidenPostgresDatabaseNames,database.Database) {
		return false
	}
	return true
}

func (database *PostgresDatabase) GetCreateQuery() (query string) {
	query = "CREATE DATABASE " + database.Database
	return
}

func (database *PostgresDatabase) GetSetPrivilegesQuery() (query string) {
	query = "GRANT ALL PRIVILEGES ON DATABASE " + database.Database + " TO " + database.Username
	return
}

func (database *PostgresDatabase) GetDeleteQuery() (query string) {
	query = "DROP DATABASE " + database.Database
	return
}
