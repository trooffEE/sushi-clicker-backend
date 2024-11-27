package main

import (
	"context"
	_ "database/sql"
	_ "github.com/lib/pq"
	"github.com/trooffEE/sushi-clicker-backend/internal/app"
	_ "github.com/trooffEE/sushi-clicker-backend/internal/config"
	"github.com/trooffEE/sushi-clicker-backend/internal/db"
	"os"
	"os/signal"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	database := db.NewDatabaseClient()
	app.InitServer(ctx, database)

	<-ctx.Done()
}
