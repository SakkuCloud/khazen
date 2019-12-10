package model

import "khazen/config"

type MySQLDatabase struct {
	Username     string `json:"username"`
	Database     string `json:"database"`
	CharacterSet string `json:"character_set"`
}

func (database *MySQLDatabase) HasRequirements() bool {
	if database.Username == "" || database.Database == "" {
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
