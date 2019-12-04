package handler

import (
	"encoding/json"
	"khazen/app/model"
	"khazen/app/service"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func CreateMySQLDatabase(w http.ResponseWriter, r *http.Request) {
	database := model.MySQLDatabase{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&database); err != nil {
		log.Warnf("Cannot decode mysql database object in CreateMySQLDatabase, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot decode object")
		return
	}
	defer r.Body.Close()

	if !database.HasRequirements() {
		log.Warnf("Invalid args in CreateMySQLDatabase request")
		respondMessage(w, http.StatusBadRequest, "Invalid args request")
		return
	}

	err := service.DatabaseExecute(database.GetCreateQuery())
	if err != nil {
		log.Warnf("Cannot create mysql database, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot create mysql database")
		return
	}

	err = service.DatabaseExecute(database.GetSetPrivilegesQuery())
	if err != nil {
		log.Warnf("Cannot set privileges in database, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot set privileges in database")
		return
	}

	log.Infof("Mysql database created, %s", database.Username)
	respondJSON(w, http.StatusCreated, database)
}
