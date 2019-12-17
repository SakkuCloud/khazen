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

func ImportMySQLDatabase(w http.ResponseWriter, r *http.Request) {
	databaseName := mux.Vars(r)["name"]
	if databaseName == "" {
		log.Warnf("Invalid request, empty mysql database name in ImportMySQLDatabase")
		respondMessage(w, http.StatusBadRequest, "Invalid request, empty mysql database name")
		return
	}

	if err := r.ParseMultipartForm(config.ImportMaxFile); err != nil {
		log.Warnf("Cannot parse body in ImportMySQLDatabase, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot parse body")
		return
	}

	file, handler, err := r.FormFile(config.ImportFileKey)
	if err != nil {
		log.Warnf("Cannot retrieving the import file in ImportMySQLDatabase, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot retrieving the import file")
		return
	}
	defer file.Close()
	defer r.Body.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Warnf("Cannot read file byte in ImportMySQLDatabase, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot read file")
		return
	}

	tempFile, err := ioutil.TempFile(config.TmpDirectory, config.ImportTmpFilePattern)
	if err != nil {
		log.Warnf("Cannot create tmp file in ImportMySQLDatabase, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot create file")
		return
	}
	defer tempFile.Close()

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		log.Warnf("Cannot write to tmp file in ImportMySQLDatabase, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot write file")
		_ = os.Remove(tempFile.Name())
		return
	}

	cmd := config.Config.MySQLCmd + " -u " + config.Config.MySQL.User + " -p" + config.Config.MySQL.Password  + " " + databaseName + " < " + tempFile.Name()
	if err := service.OSCommandExecute(cmd); err != nil {
		log.Warnf("Cannot import tmp file in ImportMySQLDatabase, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot import file")
		_ = os.Remove(tempFile.Name())
		return
	}

	_ = os.Remove(tempFile.Name())
	log.Infof("Mysql %s database imported", databaseName)
	respondJSON(w, http.StatusOK, handler.Filename)
}

func ImportPostgresDatabase(w http.ResponseWriter, r *http.Request) {
	databaseName := mux.Vars(r)["name"]
	if databaseName == "" {
		log.Warnf("Invalid request, empty postgres database name in ImportPostgresDatabase")
		respondMessage(w, http.StatusBadRequest, "Invalid request, empty postgres database name")
		return
	}

	if err := r.ParseMultipartForm(config.ImportMaxFile); err != nil {
		log.Warnf("Cannot parse body in ImportPostgresDatabase, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot parse body")
		return
	}

	file, handler, err := r.FormFile(config.ImportFileKey)
	if err != nil {
		log.Warnf("Cannot retrieving the import file in ImportPostgresDatabase, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot retrieving the import file")
		return
	}
	defer file.Close()
	defer r.Body.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Warnf("Cannot read file byte in ImportPostgresDatabase, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot read file")
		return
	}

	tempFile, err := ioutil.TempFile(config.TmpDirectory, config.ImportTmpFilePattern)
	if err != nil {
		log.Warnf("Cannot create tmp file in ImportPostgresDatabase, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot create file")
		return
	}
	defer tempFile.Close()

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		log.Warnf("Cannot write to tmp file in ImportPostgresDatabase, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot write file")
		_ = os.Remove(tempFile.Name())
		return
	}

	cmd := config.Config.PostgresCmd + " -U " + config.Config.Postgres.User + " -d " + databaseName + " -1 -f " + tempFile.Name()
	if err := service.OSCommandExecute(cmd); err != nil {
		log.Warnf("Cannot import tmp file in ImportPostgresDatabase, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot import file")
		_ = os.Remove(tempFile.Name())
		return
	}

	_ = os.Remove(tempFile.Name())
	log.Infof("Postgres %s database imported", databaseName)
	respondJSON(w, http.StatusOK, handler.Filename)
}
