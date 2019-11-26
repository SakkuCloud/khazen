package handler

import (
	"encoding/json"
	"khazen/app/model"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func CreateMySQLDatabase(w http.ResponseWriter, r *http.Request, dbURI string) {
	database := model.MySQLDatabase{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&database); err != nil {
		log.Warnf("Cannot decode mysql database object in CreateMySQLDatabase, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot decode object")
		return
	}
	defer r.Body.Close()

	err := database.Create(dbURI)
	if err != nil {
		log.Warnf("Cannot create mysql database, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot create mysql database")
		return
	}

	err = database.SetPrivileges(dbURI)
	if err != nil {
		log.Warnf("Cannot set privileges in database, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot set privileges in database")
		return
	}

	log.Infof("Mysql database created, %s", database.Username)
	respondJSON(w, http.StatusCreated, database)
}
