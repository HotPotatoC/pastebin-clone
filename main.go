package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/HotPotatoC/pastebin-clone/api"
	"github.com/HotPotatoC/pastebin-clone/backend"
	"github.com/HotPotatoC/pastebin-clone/clients"
	"github.com/HotPotatoC/pastebin-clone/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if os.Getenv("ENVIRONMENT") == "dev" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		log.Logger = log.With().Caller().Logger()
	}

	log.Info().Msg("Starting pastebin-clone backend")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	httpHostname, ok := os.LookupEnv("HOST")
	if !ok {
		httpHostname = "0.0.0.0"
	}

	httpPort, ok := os.LookupEnv("PORT")
	if !ok {
		httpPort = "9000"
	}

	scyllaKeyspace, ok := os.LookupEnv("SCYLLADB_KEYSPACE")
	if !ok {
		scyllaKeyspace = "pastebin"
	}

	var scyllaHosts []string
	scyllaHost, ok := os.LookupEnv("SCYLLADB_HOSTS")
	if !ok {
		scyllaHost = "127.0.0.1:9042"
	}

	scyllaHosts = strings.Split(scyllaHost, ",")

	log.Info().Any("hosts", scyllaHosts).Msgf("Connecting to ScyllaDB")
	db, err := clients.NewScyllaDB(ctx, scyllaKeyspace, scyllaHosts)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer db.Close()

	app := fiber.New(fiber.Config{
		AppName:      "github.com/HotPotatoC/pastebin-clone - Backend",
		BodyLimit:    1024 * 1024 * 6, // 6MB
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		ErrorHandler: api.Error,
	})

	repository := repository.Dependency{
		DB: db,
	}

	backend := backend.Dependency{
		Repository: repository,
	}

	api := api.Dependency{
		Backend: backend,
	}

	corsMiddleware := cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{fiber.MethodGet, fiber.MethodPost, fiber.MethodHead}, ","),
		AllowHeaders: fiber.HeaderAuthorization,
	})

	app.Use(corsMiddleware)

	app.Get("/health", api.Health)

	identityAPI := app.Group("/identity")       // /identity
	identityAPI.Post("/register", api.Register) // /identity/register
	identityAPI.Post("/login", api.Login)       // /identity/login

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Info().Msgf("Listening on %s", net.JoinHostPort(httpHostname, httpPort))
		err := app.Listen(net.JoinHostPort(httpHostname, httpPort))
		if err != nil {
			log.Fatal().Err(err).Send()
		}
	}()

	signal := <-exitSignal
	log.Warn().Any("signal", signal).Msg("received signal, shutting down...")
	app.Shutdown()
}
