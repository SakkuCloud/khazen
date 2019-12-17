package service

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"khazen/config"
)

func MySQLDatabaseExecute(query string) (err error) {
	var db *sql.DB
	var res sql.Result

	db, err = sql.Open(config.DataBaseDriverName, getMySQLDatabaseURI())
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

func getMySQLDatabaseURI() (dbURI string) {
	dbURI = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True",
		config.Config.MySQL.User,
		config.Config.MySQL.Password,
		config.Config.MySQL.Host,
		config.Config.MySQL.Port)
	return
}