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

	a := api.New(conf, mux)

	t := &Tracker{db: db, conf: conf, api: a}

	t.InitRoute(mux)

	return t
}

func (t *Tracker) Start() error {
	return t.api.Start()
}
