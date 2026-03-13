package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/Revan84/homeapp_backend/internal/config"
	"github.com/Revan84/homeapp_backend/internal/database"
	httpapp "github.com/Revan84/homeapp_backend/internal/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := config.Load()

	db, err := database.NewPostgresDB(cfg.DatabaseURL())
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}
	defer db.Close()

	log.Println("Database connected successfully")

	router := httpapp.NewRouter(db, cfg)

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}

	log.Printf("API listening on :%s", cfg.AppPort)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
