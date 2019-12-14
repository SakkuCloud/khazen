package handler

import (
	"encoding/json"
	"khazen/app/model"
	"khazen/app/service"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func CreateMySQLAccount(w http.ResponseWriter, r *http.Request) {
	account := model.MySQLAccount{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&account); err != nil {
		log.Warnf("Cannot decode mysql account object in CreateMySQLAccount, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot decode object")
		return
	}
	defer r.Body.Close()

	if !account.HasRequirements() {
		log.Warnf("Invalid args in CreateMySQLAccount request")
		respondMessage(w, http.StatusBadRequest, "Invalid args request")
		return
	}

	if err := service.DatabaseExecute(account.GetCreateQuery()); err != nil {
		log.Warnf("Cannot create mysql account, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot create mysql account")
		return
	}

	log.Infof("Mysql account created, %s", account.Username)
	respondJSON(w, http.StatusCreated, account)
}
