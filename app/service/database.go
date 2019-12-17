package service

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"khazen/config"
)

func MySQLDatabaseExecute(query string) (err error) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True",
		config.Config.MySQL.User,
		config.Config.MySQL.Password,
		config.Config.MySQL.Host,
		config.Config.MySQL.Port)
	err = databaseExecute("mysql", query, dbURI)
	return
}

func PostgresDatabaseExecute(query string) (err error) {
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/?sslmode=disable",
		config.Config.Postgres.User,
		config.Config.Postgres.Password,
		config.Config.Postgres.Host,
		config.Config.Postgres.Port)
	err = databaseExecute("postgres", query, dbURI)
	return
}

func databaseExecute(driver string, query string, uri string) (err error) {
	var db *sql.DB
	var res sql.Result

	db, err = sql.Open(driver, uri)
	if err == nil {
		defer db.Close()

		log.Debug(query)
		res, err = db.Exec(query)
		if res != nil {
			rowsNo, _ := res.RowsAffected()
			lastInsertId, _ := res.LastInsertId()
			log.Debugf("rows affected: %d", rowsNo)
			log.Debugf("last insert id: %d", lastInsertId)
		}
	}
	return
}
