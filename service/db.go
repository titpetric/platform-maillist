package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/titpetric/platform"
)

// DB will return the sqlx.DB in use for the maillist module.
func DB() (*sqlx.DB, error) {
	return platform.Database.Connect("maillist")
}
