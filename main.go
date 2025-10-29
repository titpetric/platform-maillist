package main

import (
	"log"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/titpetric/platform"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "modernc.org/sqlite"

	_ "github.com/titpetric/platform/module/autoload"

	"github.com/titpetric/platform-maillist/service"
)

func main() {
	// Register common middleware.
	platform.AddMiddleware(middleware.Logger)

	platform.AddModule(service.NewMailList())

	if err := start(); err != nil {
		log.Fatalf("exit error: %v", err)
	}
}

func start() error {
	if err := Migrate(); err != nil {
		return err
	}

	return platform.Start()
}
