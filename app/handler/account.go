package handler

import (
	"encoding/json"
	"khazen/app/model"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func CreateMySQLAccount(w http.ResponseWriter, r *http.Request, dbURI string) {
	account := model.MySQLAccount{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&account); err != nil {
		log.Warnf("Cannot decode mysql account object in CreateMySQLAccount, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot decode object")
		return
	}
	defer r.Body.Close()

	err := account.Create(dbURI)
	if err != nil {
		log.Warnf("Cannot create account, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot create mysql account")
	} else {
		log.Infof("Mysql account created, %s", account.Username)
		respondJSON(w, http.StatusCreated, account)
	}
}
