package main

import (
	"shared-db/api"
	"shared-db/conf"
	"shared-db/db"
	"shared-db/internal/tracker"
)

func main() {
	// 1. Get config
	c := conf.New()

	// 2. Get DB
	d := db.New(c)

	// 3. api
	a := api.New(d, c)

	// 4. init route
	a.InitRoute("/tracker", tracker.GetTracker)

	// 5. start api
	a.Start()
}
