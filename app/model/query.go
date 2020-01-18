package model

type Query struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	QueryString string `json:"query_string"`
	QueryType   int    `json:"query_type"`
}

type QueryResult struct {
	Columns      []string   `json:"columns"`
	Rows         [][]string `json:"rows"`
	RowsAffected int64      `json:"rows_affected"`
	LastInsertId int64      `json:"last_insert_id"`
}

func (query *Query) HasRequirements() bool {
	if query.Username == "" || query.Password == "" || query.QueryString == "" {
		return false
	}
	return true
}
