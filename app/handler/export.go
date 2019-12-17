package handler

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"khazen/app/service"
	"khazen/config"
	"net/http"
	"os"
)

func ExportMySQLDatabase(w http.ResponseWriter, r *http.Request) {
	databaseName := mux.Vars(r)["name"]
	if databaseName == "" {
		log.Warnf("Invalid request, empty mysql database name in ExportMySQLDatabase")
		respondMessage(w, http.StatusBadRequest, "Invalid request, empty mysql database name")
		return
	}

	tempFile, err := ioutil.TempFile(config.TmpDirectory, databaseName+"-*.sql.gz")
	if err != nil {
		log.Warnf("Cannot create tmp file in ExportMySQLDatabase, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot create file")
		return
	}
	defer tempFile.Close()

	cmd := config.Config.MySQLDumpCmd + " -q --default-character-set " + config.DefaultDatabaseCharacterSet + " -p" + config.Config.MySQL.Password
	cmd = cmd + " " + databaseName + " | gzip > " + tempFile.Name()
	if err := service.OSCommandExecute(cmd); err != nil {
		log.Warnf("Cannot export to tmp file in ExportMySQLDatabase, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot export file")
		_ = os.Remove(tempFile.Name())
		return
	}

	if config.Config.UseSakkuUploadFileService {
		sakkuUserId := r.FormValue("sakku_user_id")
		if err := service.UploadFile(config.SakkuUploadFileKeyFile, tempFile.Name(), config.SakkuUploadFileEndpoint+sakkuUserId, databaseName); err != nil {
			log.Warnf("Cannot upload file to sakku in ExportMySQLDatabase, %s", err.Error())
			respondMessage(w, http.StatusInternalServerError, "Cannot upload file to sakku")
			_ = os.Remove(tempFile.Name())
			return
		}
		_ = os.Remove(tempFile.Name())
	}

	log.Infof("Mysql %s database exported", databaseName)
	respondJSON(w, http.StatusOK, databaseName)
}

func ExportPostgresDatabase(w http.ResponseWriter, r *http.Request) {
	databaseName := mux.Vars(r)["name"]
	if databaseName == "" {
		log.Warnf("Invalid request, empty postgres database name in ExportPostgresDatabase")
		respondMessage(w, http.StatusBadRequest, "Invalid request, empty postgres database name")
		return
	}

	tempFile, err := ioutil.TempFile(config.TmpDirectory, databaseName+"-*.sql.gz")
	if err != nil {
		log.Warnf("Cannot create tmp file in ExportPostgresDatabase, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot create file")
		return
	}
	defer tempFile.Close()

	cmd := config.Config.PostgresDumpCmd + " -U " + config.Config.Postgres.User + " " + databaseName + " | gzip > " + tempFile.Name()
	if err := service.OSCommandExecute(cmd); err != nil {
		log.Warnf("Cannot export to tmp file in ExportPostgresDatabase, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot export file")
		_ = os.Remove(tempFile.Name())
		return
	}

	if config.Config.UseSakkuUploadFileService {
		sakkuUserId := r.FormValue("sakku_user_id")
		if err := service.UploadFile(config.SakkuUploadFileKeyFile, tempFile.Name(), config.SakkuUploadFileEndpoint+sakkuUserId, databaseName); err != nil {
			log.Warnf("Cannot upload file to sakku in ExportPostgresDatabase, %s", err.Error())
			respondMessage(w, http.StatusInternalServerError, "Cannot upload file to sakku")
			_ = os.Remove(tempFile.Name())
			return
		}
		_ = os.Remove(tempFile.Name())
	}

	log.Infof("Postgres %s database exported", databaseName)
	respondJSON(w, http.StatusOK, databaseName)
}
