package handler

import (
	"khazen/app/model"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func IsAuthorized(w http.ResponseWriter, r *http.Request, auth *model.Auth) bool {
	if !auth.IsAuthorized(r.Header.Get("service"), r.Header.Get("service-key")) {
		log.Warnf("Unauthorized, %s", r.RemoteAddr)
		respondMessage(w, http.StatusUnauthorized, "Unauthorized")
		return false
	}
	return true
}
