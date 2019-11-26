package handler

import (
	"encoding/json"
	"khazen/app/model"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func ExecMySQLBundle(w http.ResponseWriter, r *http.Request, dbURI string) {
	mysql := model.MySQLBundle{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&mysql); err != nil {
		log.Infof("Cannot decode mysql bundle object in ExecMySQLBundle, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot decode object")
		return
	}
	defer r.Body.Close()

	err := mysql.Account.Create(dbURI)
	if err != nil {
		log.Warnf("Cannot create account in exec mysql bundle, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot create account in exec mysql")
		return
	}

	err = mysql.Database.Create(dbURI)
	if err != nil {
		log.Warnf("Cannot create database in exec mysql bundle, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot create database in exec mysql")
		return
	}

	err = mysql.Database.SetPrivileges(dbURI)
	if err != nil {
		log.Warnf("Cannot set privileges in exec mysql bundle, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot set privileges in exec mysql bundle")
		return
	}

	log.Infof("Mysql bundle executed, %s", mysql.Account.Username)
	respondJSON(w, http.StatusCreated, mysql)
}
