package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
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

	database.SetDefaults()
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

	log.Infof("Mysql database created, %s", database.Database)
	respondJSON(w, http.StatusCreated, database)
}

func DeleteMySQLDatabase(w http.ResponseWriter, r *http.Request) {
	databaseName := mux.Vars(r)["name"]
	if databaseName == "" {
		log.Warnf("Invalid request, empty mysql database name")
		respondMessage(w, http.StatusBadRequest, "Invalid request, empty mysql database name")
		return
	}

	database := model.MySQLDatabase{Database: databaseName}
	err := service.DatabaseExecute(database.GetDeleteQuery())
	if err != nil {
		log.Warnf("Cannot delete mysql database, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot delete mysql database")
		return
	}

	log.Infof("Mysql database deleted, %s", database.Database)
	respondJSON(w, http.StatusNoContent, nil)
}
