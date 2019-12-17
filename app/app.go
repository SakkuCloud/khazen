package app

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"khazen/app/handler"
	"khazen/app/model"
	"khazen/config"
	"net/http"
	"os"
)

type App struct {
	Router   *mux.Router
	Auth     *model.Auth
}

// INIT
func (a *App) Init() {
	a.Router = mux.NewRouter()
	a.setRouters()

	a.Auth = &model.Auth{
		AccessKey: config.Config.AccessKey,
		SecretKey: config.Config.SecretKey,
	}

	_ = os.Mkdir(config.TmpDirectory, os.ModeDir)
}

func (a *App) setRouters() {
	APISubRouter := a.Router.PathPrefix("/api").Subrouter()

	APISubRouter.HandleFunc("/health", a.GetHealth).Methods(http.MethodGet)

	APISubRouter.HandleFunc("/mysql/bundle", a.ExecMySQLBundle).Methods(http.MethodPost)
	APISubRouter.HandleFunc("/mysql/account", a.CreateMySQLAccount).Methods(http.MethodPost)
	APISubRouter.HandleFunc("/mysql/database", a.CreateMySQLDatabase).Methods(http.MethodPost)
	APISubRouter.HandleFunc("/mysql/database/{name}", a.DeleteMySQLDatabase).Methods(http.MethodDelete)
	APISubRouter.HandleFunc("/mysql/import/{name}", a.ImportMySQLDatabase).Methods(http.MethodPost)
	APISubRouter.HandleFunc("/mysql/export/{name}", a.ExportMySQLDatabase).Methods(http.MethodGet)

	APISubRouter.HandleFunc("/postgres/bundle", a.ExecPostgresBundle).Methods(http.MethodPost)
	APISubRouter.HandleFunc("/postgres/account", a.CreatePostgresAccount).Methods(http.MethodPost)
	APISubRouter.HandleFunc("/postgres/database", a.CreatePostgresDatabase).Methods(http.MethodPost)
	APISubRouter.HandleFunc("/postgres/database/{name}", a.DeletePostgresDatabase).Methods(http.MethodDelete)
}

// HEALTH
func (a *App) GetHealth(w http.ResponseWriter, r *http.Request) {
	handler.GetHealth(w, r)
}

// MYSQL
func (a *App) CreateMySQLAccount(w http.ResponseWriter, r *http.Request) {
	if handler.IsAuthorized(w, r, a.Auth) {
		handler.CreateMySQLAccount(w, r)
	}
}
func (a *App) CreateMySQLDatabase(w http.ResponseWriter, r *http.Request) {
	if handler.IsAuthorized(w, r, a.Auth) {
		handler.CreateMySQLDatabase(w, r)
	}
}
func (a *App) DeleteMySQLDatabase(w http.ResponseWriter, r *http.Request) {
	if handler.IsAuthorized(w, r, a.Auth) {
		handler.DeleteMySQLDatabase(w, r)
	}
}
func (a *App) ExecMySQLBundle(w http.ResponseWriter, r *http.Request) {
	if handler.IsAuthorized(w, r, a.Auth) {
		handler.ExecMySQLBundle(w, r)
	}
}
func (a *App) ImportMySQLDatabase(w http.ResponseWriter, r *http.Request) {
	if handler.IsAuthorized(w, r, a.Auth) {
		handler.ImportMySQLDatabase(w, r)
	}
}
func (a *App) ExportMySQLDatabase(w http.ResponseWriter, r *http.Request) {
	if handler.IsAuthorized(w, r, a.Auth) {
		handler.ExportMySQLDatabase(w, r)
	}
}

// Postgres
func (a *App) CreatePostgresAccount(w http.ResponseWriter, r *http.Request) {
	if handler.IsAuthorized(w, r, a.Auth) {
		handler.CreatePostgresAccount(w, r)
	}
}
func (a *App) CreatePostgresDatabase(w http.ResponseWriter, r *http.Request) {
	if handler.IsAuthorized(w, r, a.Auth) {
		handler.CreatePostgresDatabase(w, r)
	}
}
func (a *App) DeletePostgresDatabase(w http.ResponseWriter, r *http.Request) {
	if handler.IsAuthorized(w, r, a.Auth) {
		handler.DeletePostgresDatabase(w, r)
	}
}
func (a *App) ExecPostgresBundle(w http.ResponseWriter, r *http.Request) {
	if handler.IsAuthorized(w, r, a.Auth) {
		handler.ExecPostgresBundle(w, r)
	}
}

// RUN
func (a *App) Run(host string) {
	log.Infoln("Starting ...")
	log.Fatal(http.ListenAndServe(host, a.Router))
}
