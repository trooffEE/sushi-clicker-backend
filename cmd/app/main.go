package main

import (
	_ "database/sql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/trooffEE/sushi-clicker-backend/internal/config"
	"github.com/trooffEE/sushi-clicker-backend/internal/db"
	"github.com/trooffEE/sushi-clicker-backend/internal/db/repository"
	"github.com/trooffEE/sushi-clicker-backend/internal/handlers/auth"
	"github.com/trooffEE/sushi-clicker-backend/internal/service/user"
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
	usrRepo := repository.NewUserRepository(s.DB)
	usrService := user.NewUserService(usrRepo)

	hAuth := auth.NewHandler(usrService)

	s.Router.HandleFunc("/api/auth/login", hAuth.Login).Methods("POST")
	s.Router.HandleFunc("/api/auth/register", hAuth.Register).Methods("POST")
}

func (s *Server) Start() {
	srv := &http.Server{
		Handler: handlers.CORS(
			handlers.AllowedOrigins([]string{"http://localhost:5173"}), // TODO local development, later will add normal domain
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
