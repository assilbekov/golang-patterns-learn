package db

import (
	"fmt"
	"log"
	"shared-db/conf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
		db *gorm.DB
}

func New(conf *conf.Config) *DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=test port=%s sslmode=disable",
		conf.DB_HOST,
		conf.DB_USER,
		conf.DB_PASSWORD,
		conf.DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		db: db,
	}
}


