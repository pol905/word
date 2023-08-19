package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pol905/word/entities"
	"github.com/pol905/word/handlers"
	customMiddleware "github.com/pol905/word/middleware"
	"github.com/pol905/word/repositories"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimestampFieldName = "ts"
	zerolog.LevelFieldName = "lvl"
	zerolog.MessageFieldName = "msg"
	zerolog.DurationFieldUnit = time.Millisecond

	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	dsn := "postgres://postgres:postgres@localhost:5432/word?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("Failed to connect to database")
	}

	db.AutoMigrate(entities.Book{})

	router := chi.NewRouter()

	router.Use(
		middleware.Heartbeat("/healthz"),
		middleware.RealIP,
		middleware.SetHeader("Content-Type", "application/json"),
		middleware.Compress(5, "gzip"),
		customMiddleware.Logger(&log),
		middleware.Timeout(2*time.Second),
	)

	bookRepository := repositories.NewBookRepository(db)
	bookRouter := handlers.BookRouter(bookRepository)

	router.Mount("/books", bookRouter)

	log.Info().Msg("Starting Server on port 8080")
	log.Fatal().Stack().Err(http.ListenAndServe(":8080", router)).Msg("")
}
