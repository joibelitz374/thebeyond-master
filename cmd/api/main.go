package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/quickpowered/thebeyond-master/cmd/api/internal/handlers"
	"github.com/quickpowered/thebeyond-master/cmd/api/internal/middlewares"
	"github.com/quickpowered/thebeyond-master/internal/repositories/db"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager"
	"github.com/quickpowered/thebeyond-master/internal/services/account"
	"github.com/quickpowered/thebeyond-master/pkg/logger"
	"github.com/quickpowered/thebeyond-master/pkg/postgres"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(logger.LogLevel())

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	pgPool, err := postgres.New(ctx, postgres.DSN())
	if err != nil {
		logger.Error("failed to connect to database", zap.Error(err))
		return
	}

	accountDB := db.NewAccount(pgPool)
	accountService := account.NewService(accountDB)

	clustersPath := os.Getenv("CLUSTERS_PATH")
	if clustersPath != "" {
		clustersPath = "/etc/thebeyond/xray/clusters.yaml"
	}

	xraymanagerRepo, err := xraymanager.New(clustersPath)
	if err != nil {
		logger.Error("failed to load clusters", zap.Error(err))
		return
	}

	handlers := handlers.NewSubscription(accountService, xraymanagerRepo)
	authMiddleware := middlewares.NewAuth(os.Getenv("TG_TOKEN"))

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	app.Use(fiberLogger.New(fiberLogger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	sub := app.Group("/sub")
	sub.Get("/:sub_id", handlers.Default)
	sub.Get("/:sub_id/smart/:region", handlers.Smart)

	api := app.Group("/api")
	api.Use(authMiddleware.Init())
	api.Get("/auth", handlers.Auth)
	api.Get("/", handlers.Get)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	certFile := os.Getenv("CERT_FILE")
	if certFile == "" {
		certFile = "/etc/ssl/certs/certificate.crt"
	}

	certKeyFile := os.Getenv("CERT_KEY_FILE")
	if certKeyFile == "" {
		certKeyFile = "/etc/ssl/private/certificate.key"
	}

	log.Fatalln(app.Listen(":"+port, fiber.ListenConfig{
		CertFile:    certFile,
		CertKeyFile: certKeyFile,
	}))
}
