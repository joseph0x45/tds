package main

import (
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"server/internal/repository/postgres"
	"server/internal/rest"
	m "server/internal/rest/middleware"
	"server/services/device"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
}

func main() {
	dbURL := os.Getenv("POSTGRES_URL")
	if dbURL == "" {
		panic("DB_URL environment variable not set")
	}
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT environment variable not set")
	}
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	log.Println("Connected to database")
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	//setup router
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//repositories
	deviceRepo := postgres.NewDeviceRepo(db)
	adminRepo := postgres.NewAdminRepo(db)
	sessionRepo := postgres.NewSessionRepo(db)

	//services
	deviceService := device.NewService(logger, deviceRepo)

	//middlewares
	authMiddleware := m.NewAuthorizationMiddleware(adminRepo, sessionRepo, logger)

	//handlers
	rest.NewDeviceHandler(r, deviceService, authMiddleware)

	//launch HTTP server
	server := http.Server{
		Addr:         net.JoinHostPort("0.0.0.0", port),
		Handler:      r,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	log.Println("Server started on port", port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
