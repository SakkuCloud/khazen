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