package main

import (
	_ "database/sql"
	_ "github.com/lib/pq"
	"github.com/trooffEE/sushi-clicker-backend/internal/app"
	_ "github.com/trooffEE/sushi-clicker-backend/internal/config"
	"github.com/trooffEE/sushi-clicker-backend/internal/db"
)

func main() {
	database := db.NewDatabaseClient()
	app.InitServer(database)
}
