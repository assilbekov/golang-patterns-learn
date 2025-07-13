package user

import (
	"net/http"
	"shared-db/api"
	"shared-db/conf"
	"shared-db/db"
)

type User struct {
	db *db.DB
	conf *conf.Config
	api *api.API
}

func New(db *db.DB, conf *conf.Config) *User {
	mux := http.NewServeMux()

	a := api.New(conf, mux)

	t := &User{db: db, conf: conf, api: a}

	t.InitRoute(mux)

	return t
}

func (t *User) Start() error {
	return t.api.Start()
}
