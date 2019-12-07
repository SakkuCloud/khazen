package handler

import (
	log "github.com/sirupsen/logrus"
	"khazen/app/model"
	"net/http"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	health := model.Health{}

	//todo: Check Endpoints and services health!
	health.Status = true

	health.SetUptime()
	health.SetServerTime()

	log.Debugf("Health sent!, health: %s", health.Status)
	respondJSON(w, http.StatusOK, health)
}