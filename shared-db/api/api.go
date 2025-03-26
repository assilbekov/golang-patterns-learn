package api

import (
	"fmt"
	"log"
	"net/http"
	"shared-db/conf"
	"shared-db/db"
)

type API struct {
	apiPort string
	db *db.DB
	server *http.Server
}

func New(db *db.DB, conf *conf.Config) *API {
	//mux := http.NewServeMux()
	return &API{
		apiPort: conf.API_PORT,
		db: db,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%s", conf.API_PORT),
			//Handler: mux,
		},
	}
}

func (a *API) Start() error {
	log.Printf("Starting HTTP server on %s", a.apiPort)
	return a.server.ListenAndServe()
}

func (a *API) InitRoute(path string, handler http.HandlerFunc) {
	a.server.Handler = http.HandlerFunc(path, handler)
}
