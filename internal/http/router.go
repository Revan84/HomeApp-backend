package httpapp

import (
	"database/sql"
	"net/http"

	"github.com/Revan84/homeapp_backend/internal/auth"
	"github.com/Revan84/homeapp_backend/internal/config"
)

func NewRouter(db *sql.DB, cfg config.Config) http.Handler {

	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Auth module
	authRepo := auth.NewRepository(db)

	jwtManager := &auth.JWTManager{
		Secret: cfg.JWTSecret,
	}

	authService := auth.NewService(authRepo, jwtManager)

	authHandler := auth.NewHandler(authService)

	mux.HandleFunc("/api/v1/auth/register", authHandler.Register)
	mux.HandleFunc("/api/v1/auth/login", authHandler.Login)

	return mux
}
