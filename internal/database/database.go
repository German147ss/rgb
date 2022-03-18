package database

import (
	"github.com/go-pg/pg/v10"
)

func NewDBOptions() *pg.Options {
	return &pg.Options{
		Addr:     "localhost:5433",
		Database: "rgb",
		User:     "postgres",
		Password: "new_password",
	}
}
