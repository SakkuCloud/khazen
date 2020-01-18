package service

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"khazen/app/model"
	"khazen/config"
)

func MySQLDatabaseExecute(query string, db string) (queryResult *model.QueryResult, err error) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		config.Config.MySQL.User,
		config.Config.MySQL.Password,
		config.Config.MySQL.Host,
		config.Config.MySQL.Port,
		db)
	queryResult, err = databaseExecute("mysql", query, dbURI)
	return
}

func MySQLDatabaseQuery(query string, db string) (queryResult *model.QueryResult, err error) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		config.Config.MySQL.User,
		config.Config.MySQL.Password,
		config.Config.MySQL.Host,
		config.Config.MySQL.Port,
		db)
	queryResult, err = databaseQuery("mysql", query, dbURI)
	return
}

func PostgresDatabaseExecute(query string, db string) (queryResult *model.QueryResult, err error) {
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Config.Postgres.User,
		config.Config.Postgres.Password,
		config.Config.Postgres.Host,
		config.Config.Postgres.Port,
		db)
	queryResult, err = databaseExecute("postgres", query, dbURI)
	return
}

func PostgresDatabaseQuery(query string, db string) (queryResult *model.QueryResult, err error) {
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Config.Postgres.User,
		config.Config.Postgres.Password,
		config.Config.Postgres.Host,
		config.Config.Postgres.Port,
		db)
	queryResult, err = databaseQuery("postgres", query, dbURI)
	return
}

func databaseExecute(driver string, query string, uri string) (queryResult *model.QueryResult, err error) {
	var db *sql.DB
	var res sql.Result
	queryResult = &model.QueryResult{}
	log.Debugf("Database execute in %s = %s",driver,query)

	db, err = sql.Open(driver, uri)
	if err == nil {
		defer db.Close()
		res, err = db.Exec(query)
		if res != nil {
			queryResult.RowsAffected, _ = res.RowsAffected()
			queryResult.LastInsertId, _ = res.LastInsertId()
			log.Debugf("rows affected: %d", queryResult.RowsAffected)
			log.Debugf("last insert id: %d", queryResult.LastInsertId)
		}
	}
	return
}

func databaseQuery(driver string, query string, uri string) (queryResult *model.QueryResult, err error) {
	var db *sql.DB
	var rows *sql.Rows
	queryResult = &model.QueryResult{}
	log.Debugf("Database query in %s = %s",driver,query)

	db, err = sql.Open(driver, uri)
	if err == nil {
		defer db.Close()
		rows, err = db.Query(query)
		if err == nil {
			queryResult.Columns, err = rows.Columns()
			if err == nil {
				for rows.Next() {
					readCols := make([]interface{}, len(queryResult.Columns))
					row := make([]string, len(queryResult.Columns))
					for i, _ := range row {
						readCols[i] = &row[i]
					}

					err := rows.Scan(readCols...)
					if err != nil {
						break
					}
					log.Debug(row)

					queryResult.Rows = append(queryResult.Rows,row)
				}
			}
		}
	}
	return
}