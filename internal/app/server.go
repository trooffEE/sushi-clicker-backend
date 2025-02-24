package app

import (
	"context"
	"errors"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/trooffEE/sushi-clicker-backend/internal/db/repository"
	authHandler "github.com/trooffEE/sushi-clicker-backend/internal/handlers/auth"
	userHandler "github.com/trooffEE/sushi-clicker-backend/internal/handlers/user"
	"github.com/trooffEE/sushi-clicker-backend/internal/middlewares"
	"github.com/trooffEE/sushi-clicker-backend/internal/service/user"
	"go.uber.org/zap"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	Router *mux.Router
	DB     *sqlx.DB
	server *http.Server
}

// Returns shutdowner
func InitServer(db *sqlx.DB) func() {
	server := &Server{
		Router: mux.NewRouter(),
		DB:     db,
	}

	server.MountMiddlewares()
	server.MountHandlers()

	return server.Start()
}

func (s *Server) ShutdownHTTPServer(ctx context.Context) {
	zap.L().Info("Shutting down HTTP server... ⚒")
	err := s.server.Shutdown(ctx)
	if err != nil {
		zap.L().Fatal("Error shutting down HTTP server", zap.Error(err))
	}

	err = s.DB.Close()
	if err != nil {
		zap.L().Fatal("Error shutting down HTTP server", zap.Error(err))
	}
	zap.L().Info("Server is down 🫂")
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

func (s *Server) Start() func() {
	srv := &http.Server{
		Handler: handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedOrigins([]string{"http://localhost:5173"}), // TODO local development, later will add normal origin
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "X-Requested-With"}),
		)(s.Router),
		Addr:         ":3010", // TODO env
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	s.server = srv

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			zap.L().Fatal("http server failed to start", zap.Error(err))
		}
	}()

	return func() {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		s.ShutdownHTTPServer(ctx)
		wg.Wait()
	}
}
