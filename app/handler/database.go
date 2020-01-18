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
	if _, err := service.MySQLDatabaseExecute(database.GetCreateQuery(),""); err != nil {
		log.Warnf("Cannot create mysql database, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot create mysql database")
		return
	}

	if _, err := service.MySQLDatabaseExecute(database.GetSetPrivilegesQuery(),""); err != nil {
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
		log.Warnf("Invalid request, empty mysql database name in DeleteMySQLDatabase")
		respondMessage(w, http.StatusBadRequest, "Invalid request, empty mysql database name")
		return
	}

	database := model.MySQLDatabase{Database: databaseName}
	if _, err := service.MySQLDatabaseExecute(database.GetDeleteQuery(),""); err != nil {
		log.Warnf("Cannot delete mysql database, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot delete mysql database")
		return
	}

	log.Infof("Mysql database deleted, %s", database.Database)
	respondJSON(w, http.StatusNoContent, nil)
}

func CreatePostgresDatabase(w http.ResponseWriter, r *http.Request) {
	database := model.PostgresDatabase{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&database); err != nil {
		log.Warnf("Cannot decode postgres database object in CreatePostgresDatabase, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot decode object")
		return
	}
	defer r.Body.Close()

	if !database.HasRequirements() {
		log.Warnf("Invalid args in CreatePostgresDatabase request")
		respondMessage(w, http.StatusBadRequest, "Invalid args request")
		return
	}

	if _, err := service.PostgresDatabaseExecute(database.GetCreateQuery(),""); err != nil {
		log.Warnf("Cannot create postgres database, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot create postgres database")
		return
	}

	if _, err := service.PostgresDatabaseExecute(database.GetSetPrivilegesQuery(),""); err != nil {
		log.Warnf("Cannot set privileges in database, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot set privileges in database")
		return
	}

	log.Infof("Postgres database created, %s", database.Database)
	respondJSON(w, http.StatusCreated, database)
}

func DeletePostgresDatabase(w http.ResponseWriter, r *http.Request) {
	databaseName := mux.Vars(r)["name"]
	if databaseName == "" {
		log.Warnf("Invalid request, empty postgres database name in DeletePostgresDatabase")
		respondMessage(w, http.StatusBadRequest, "Invalid request, empty postgres database name")
		return
	}

	database := model.PostgresDatabase{Database: databaseName}
	if _, err := service.PostgresDatabaseExecute(database.GetDeleteQuery(),""); err != nil {
		log.Warnf("Cannot delete postgres database, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot delete postgres database")
		return
	}

	log.Infof("Postgres database deleted, %s", database.Database)
	respondJSON(w, http.StatusNoContent, nil)
}