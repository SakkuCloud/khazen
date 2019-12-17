package handler

import (
	"encoding/json"
	"khazen/app/model"
	"khazen/app/service"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func ExecMySQLBundle(w http.ResponseWriter, r *http.Request) {
	mysql := model.MySQLBundle{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&mysql); err != nil {
		log.Infof("Cannot decode mysql bundle object in ExecMySQLBundle, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot decode object")
		return
	}
	defer r.Body.Close()

	if !mysql.HasRequirements() {
		log.Warnf("Invalid args in ExecMySQLBundle request")
		respondMessage(w, http.StatusBadRequest, "Invalid args request")
		return
	}

	if err := service.MySQLDatabaseExecute(mysql.Account.GetCreateQuery()); err != nil {
		log.Warnf("Cannot create account in exec mysql bundle, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot create account in exec mysql bundle")
		return
	}

	mysql.Database.SetDefaults()
	if err := service.MySQLDatabaseExecute(mysql.Database.GetCreateQuery()); err != nil {
		log.Warnf("Cannot create database in exec mysql bundle, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot create database in exec mysql bundle")
		return
	}

	if err := service.MySQLDatabaseExecute(mysql.Database.GetSetPrivilegesQuery()); err != nil {
		log.Warnf("Cannot set privileges in exec mysql bundle, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot set privileges in exec mysql bundle")
		return
	}

	log.Infof("Mysql bundle executed, %s", mysql.Account.Username)
	respondJSON(w, http.StatusCreated, mysql)
}

func ExecPostgresBundle(w http.ResponseWriter, r *http.Request) {
	postgres := model.PostgresBundle{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&postgres); err != nil {
		log.Infof("Cannot decode postgres bundle object in ExecPostgresBundle, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot decode object")
		return
	}
	defer r.Body.Close()

	if !postgres.HasRequirements() {
		log.Warnf("Invalid args in ExecPostgresBundle request")
		respondMessage(w, http.StatusBadRequest, "Invalid args request")
		return
	}

	if err := service.PostgresDatabaseExecute(postgres.Account.GetCreateQuery()); err != nil {
		log.Warnf("Cannot create account in exec postgres bundle, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot create account in exec postgres bundle")
		return
	}

	if err := service.PostgresDatabaseExecute(postgres.Database.GetCreateQuery()); err != nil {
		log.Warnf("Cannot create database in exec postgres bundle, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot create database in exec postgres bundle")
		return
	}

	if err := service.PostgresDatabaseExecute(postgres.Database.GetSetPrivilegesQuery()); err != nil {
		log.Warnf("Cannot set privileges in exec postgres bundle, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot set privileges in exec postgres bundle")
		return
	}

	log.Infof("Postgres bundle executed, %s", postgres.Account.Username)
	respondJSON(w, http.StatusCreated, postgres)
}
