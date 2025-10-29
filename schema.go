package main

import (
	"embed"
	"io/fs"
	"path"

	"github.com/go-bridget/mig/migrate"
	"github.com/titpetric/platform-maillist/service"
)

//go:embed schema
var schema embed.FS

func Migrate() error {
	db, err := service.DB()
	if err != nil {
		return err
	}

	migrations := loadMigrations(schema)

	return migrate.RunWithFS(
		db,
		migrations,
		&migrate.Options{
			Project: "maillist",
			Apply:   true,
		},
	)
}

func loadMigrations(filesystem fs.FS) migrate.FS {
	entries, err := fs.Glob(schema, "schema/*.sql")
	if err != nil {
		panic(err)
	}

	migrations := make(map[string][]byte, len(entries))

	for _, name := range entries {
		contents, _ := schema.ReadFile(name)
		migrations[path.Base(name)] = contents
	}

	return migrations
}
