package main

import (
	"shared-db/conf"
	"shared-db/db"
	"shared-db/internal/tracker"
)

func main() {
	// 1. Get config
	c := conf.New()

	// 2. Get DB
	d := db.New(c)

	// 3. Apply migrations
	if err := d.Migrate(); err != nil {
		panic("failed to run migrations: " + err.Error())
	}

	t := tracker.New(d, c)

	// 4. start api
	t.Start()
}
