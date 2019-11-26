package model

type MySQLBundle struct {
	Account  *MySQLAccount  `json:"account"`
	Database *MySQLDatabase `json:"database"`
}
