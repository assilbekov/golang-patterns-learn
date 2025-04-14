package tracker

import (
	"net/http"
	"shared-db/api"
	"shared-db/conf"
	"shared-db/db"
)

type Tracker struct {
	db *db.DB
	conf *conf.Config
	api *api.API
}

func New(db *db.DB, conf *conf.Config) *Tracker {
	mux := http.NewServeMux()

	// 4. api
	a := api.New(db, conf, mux)


	// 5. init route
	tracker.InitRoute(mux)

	return &Tracker{db: db, conf: conf, api: a}
}