package main

import (
	"context"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log"
	"movie-project/internal/app/config"
	"movie-project/internal/app/repository/repositories"
	"movie-project/internal/app/services"
	"movie-project/internal/app/transport/httpserver"
	"movie-project/internal/pkg/pg"
	"net/http"
	"time"
)

func main() {
	var pool *pgxpool.Pool

	time.Sleep(1 * time.Second)

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("config loading failed: %w", err)
	}

	// Creating pool of connections to db
	pool, err = pg.GetDB(context.Background(), cfg.DBconfig)
	if err != nil {
		log.Fatal("Creating pool of connections to db failed: %w", err)
	}

	defer pool.Close()

	// running db migrations
	if pool != nil {
		log.Println("Running PostgreSQL migrations")
		db := stdlib.OpenDBFromPool(pool)
		if err := runPGMigrations(db, cfg.MigrationsPath); err != nil {
			log.Fatal("runPgMigrations failed: %w", err)
		}
	}

	log.Println("Migrations applied successfully!")

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("IMDB Api v.1.0"))
	}).Methods(http.MethodGet)

	userRepo := repositories.NewUserRepo(pool)

	userService := services.NewUserService(userRepo)

	server := httpserver.NewAppServer(userService)

	router.HandleFunc("/user", server.CreateUser).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:    cfg.ServerConfig.Address,
		Handler: router,
	}

	log.Println("starting server on", srv.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}
	log.Printf("server listening at %s", srv.Addr)
}

func runPGMigrations(db *sql.DB, path string) error {
	if err := goose.Up(db, path); err != nil {
		log.Fatalf("Could not apply migrations: %v", err)
		return err
	}

	return nil
}
