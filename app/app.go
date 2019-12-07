package app

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"khazen/app/handler"
	"khazen/app/model"
	"khazen/config"
	"net/http"
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
}

func (a *App) setRouters() {
	APISubRouter := a.Router.PathPrefix("/api").Subrouter()

	APISubRouter.HandleFunc("/health", a.GetHealth).Methods("GET")

	APISubRouter.HandleFunc("/mysql/bundle", a.ExecMySQLBundle).Methods("POST")
	APISubRouter.HandleFunc("/mysql/account", a.CreateMySQLAccount).Methods("POST")
	APISubRouter.HandleFunc("/mysql/database", a.CreateMySQLDatabase).Methods("POST")
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

func (a *App) ExecMySQLBundle(w http.ResponseWriter, r *http.Request) {
	if handler.IsAuthorized(w, r, a.Auth) {
		handler.ExecMySQLBundle(w, r)
	}
}

// RUN
func (a *App) Run(host string) {
	log.Infoln("Starting ...")
	log.Fatal(http.ListenAndServe(host, a.Router))
}
