package app

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/trooffEE/sushi-clicker-backend/internal/db/repository"
	authHandler "github.com/trooffEE/sushi-clicker-backend/internal/handlers/auth"
	userHandler "github.com/trooffEE/sushi-clicker-backend/internal/handlers/user"
	"github.com/trooffEE/sushi-clicker-backend/internal/middlewares"
	"github.com/trooffEE/sushi-clicker-backend/internal/service/user"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Router *mux.Router
	DB     *sqlx.DB
	server *http.Server
}

func InitServer(ctx context.Context, db *sqlx.DB) {
	server := &Server{
		Router: mux.NewRouter(),
		DB:     db,
	}

	server.MountMiddlewares()
	server.MountHandlers()
	go server.Start(ctx)

	select {
	case <-ctx.Done():
		server.ShutdownHTTPServer()
	}
}

func (s *Server) ShutdownHTTPServer() {
	fmt.Println("Server shutting down... ⚒")
	err := s.server.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = s.DB.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is down 🫂")
}

func (s *Server) MountMiddlewares() {
	s.Router.Use(middlewares.HTTPHeadersMiddleware)
	s.Router.Use(middlewares.AuthMiddleware)
}

func (s *Server) MountHandlers() {
	usrRepo := repository.NewUserRepository(s.DB)
	usrService := user.NewUserService(usrRepo)

	hAuth := authHandler.NewHandler(usrService)
	hUser := userHandler.NewHandler(usrService)

	s.Router.HandleFunc("/api/auth/login", hAuth.Login).Methods("POST")
	s.Router.HandleFunc("/api/auth/register", hAuth.Register).Methods("POST")
	s.Router.HandleFunc("/api/auth/refresh-token", hAuth.RefreshToken).Methods("GET")
	s.Router.HandleFunc("/api/private/test", hUser.Test).Methods("GET")
}

func (s *Server) Start(ctx context.Context) {
	srv := &http.Server{
		Handler: handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedOrigins([]string{"http://localhost:5173"}), // TODO local development, later will add normal origin
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "X-Requested-With"}),
		)(s.Router),
		Addr:         ":3010",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	s.server = srv

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
