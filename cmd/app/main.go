package main

import (
	_ "database/sql"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/trooffEE/sushi-clicker-backend/internal/db"
	appHandlers "github.com/trooffEE/sushi-clicker-backend/internal/handlers"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Router *mux.Router
	DB     *sqlx.DB
}

func main() {
	database := db.NewDatabaseClient()
	if database == nil {
		fmt.Println("Database connection failed")
		return
	}
	server := CreateServer(database)
	server.MountHandlers()
	server.Start()
}

func CreateServer(db *sqlx.DB) *Server {
	server := &Server{
		Router: mux.NewRouter(),
		DB:     db,
	}
	return server
}

func (s *Server) MountHandlers() {
	s.Router.HandleFunc("/api/login", appHandlers.Login).Methods("POST")
}

func (s *Server) Start() {
	srv := &http.Server{
		Handler: handlers.CORS(
			handlers.AllowedOrigins([]string{"http://localhost:5173"}), // TODO local development
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "X-Requested-With"}),
		)(s.Router),
		Addr:         ":3010",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
