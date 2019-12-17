package model

type MySQLBundle struct {
	Account  *MySQLAccount  `json:"account"`
	Database *MySQLDatabase `json:"database"`
}

func (bundle *MySQLBundle) HasRequirements() bool {
	if !bundle.Account.HasRequirements() || !bundle.Database.HasRequirements() {
		return false
	}
	return true
}

type PostgresBundle struct {
	Account  *PostgresAccount  `json:"account"`
	Database *PostgresDatabase `json:"database"`
}

func (bundle *PostgresBundle) HasRequirements() bool {
	if !bundle.Account.HasRequirements() || !bundle.Database.HasRequirements() {
		return false
	}
	return true
}