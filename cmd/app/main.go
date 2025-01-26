package main

import (
	"context"
	_ "database/sql"
	_ "github.com/lib/pq"
	"github.com/trooffEE/sushi-clicker-backend/internal/app"
	"github.com/trooffEE/sushi-clicker-backend/internal/config"
	_ "github.com/trooffEE/sushi-clicker-backend/internal/config"
	"github.com/trooffEE/sushi-clicker-backend/internal/db"
	"os"
	"os/signal"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	cfg := config.NewApplicationConfig()
	database := db.NewDatabaseClient(cfg.Database)
	httpServerShutdown := app.InitServer(database)

	<-ctx.Done()
	httpServerShutdown()
}
