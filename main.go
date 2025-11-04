package main

import (
	"context"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/titpetric/platform"
	"github.com/titpetric/platform/cmd"
	"github.com/titpetric/platform/pkg/telemetry"

	_ "github.com/titpetric/platform-app/autoload"

	"github.com/titpetric/platform-maillist/service"
)

func main() {
	ctx := context.Background()

	platform.Use(middleware.Logger)
	platform.Register(service.NewMailList())

	if err := Migrate(); err != nil {
		telemetry.Fatal(ctx, err)
	}

	cmd.Main(ctx)
}
