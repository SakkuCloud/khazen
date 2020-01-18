package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"khazen/app/model"
	"khazen/app/service"
	"khazen/config"
	"net/http"
)

func QueryMySQLDatabase(w http.ResponseWriter, r *http.Request) {
	databaseName := mux.Vars(r)["name"]
	if databaseName == "" {
		log.Warnf("Invalid request, empty mysql database name in QueryMySQLDatabase")
		respondMessage(w, http.StatusBadRequest, "Invalid request, empty mysql database name")
		return
	}

	query := model.Query{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&query); err != nil {
		log.Warnf("Cannot decode mysql query object in QueryMySQLDatabase, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot decode object")
		return
	}
	defer r.Body.Close()

	if !query.HasRequirements() {
		log.Warnf("Invalid args in QueryMySQLDatabase request")
		respondMessage(w, http.StatusBadRequest, "Invalid args request")
		return
	}

	if query.QueryType == config.QueryTypeSelect {
		result, err := service.MySQLDatabaseQuery(query.QueryString,databaseName)
		if err != nil {
			log.Warnf("Cannot query in database, %s", err.Error())
			respondMessage(w, http.StatusBadRequest, "Cannot query in database")
			return
		}

		log.Infof("Select query executed in mysql database, %s", databaseName)
		respondJSON(w, http.StatusCreated, result)
	} else if query.QueryType == config.QueryTypeNonSelect {
		result, err := service.MySQLDatabaseExecute(query.QueryString,databaseName)
		if err != nil {
			log.Warnf("Cannot execute in database, %s", err.Error())
			respondMessage(w, http.StatusBadRequest, "Cannot execute in database")
			return
		}

		log.Infof("Non select query executed in mysql database, %s", databaseName)
		respondJSON(w, http.StatusCreated, result)
	} else {
		log.Warnf("Invalid query type in QueryMySQLDatabase request")
		respondMessage(w, http.StatusBadRequest, "Invalid query type")
	}
}

func QueryPostgresDatabase(w http.ResponseWriter, r *http.Request) {
	databaseName := mux.Vars(r)["name"]
	if databaseName == "" {
		log.Warnf("Invalid request, empty postgres database name in QueryPostgresDatabase")
		respondMessage(w, http.StatusBadRequest, "Invalid request, empty postgres database name")
		return
	}

	query := model.Query{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&query); err != nil {
		log.Warnf("Cannot decode postgres query object in QueryPostgresDatabase, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot decode object")
		return
	}
	defer r.Body.Close()

	if !query.HasRequirements() {
		log.Warnf("Invalid args in QueryPostgresDatabase request")
		respondMessage(w, http.StatusBadRequest, "Invalid args request")
		return
	}

	if query.QueryType == config.QueryTypeSelect {
		result, err := service.PostgresDatabaseQuery(query.QueryString,databaseName)
		if err != nil {
			log.Warnf("Cannot query in database, %s", err.Error())
			respondMessage(w, http.StatusBadRequest, "Cannot query in database")
			return
		}

		log.Infof("Select query executed in postgres database, %s", databaseName)
		respondJSON(w, http.StatusCreated, result)
	} else if query.QueryType == config.QueryTypeNonSelect {
		result, err := service.PostgresDatabaseExecute(query.QueryString,databaseName)
		if err != nil {
			log.Warnf("Cannot execute in database, %s", err.Error())
			respondMessage(w, http.StatusBadRequest, "Cannot execute in database")
			return
		}

		log.Infof("Non select query executed in postgres database, %s", databaseName)
		respondJSON(w, http.StatusCreated, result)
	} else {
		log.Warnf("Invalid query type in QueryPostgresDatabase request")
		respondMessage(w, http.StatusBadRequest, "Invalid query type")
	}
}