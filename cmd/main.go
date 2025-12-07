package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/quickpowered/frilly/internal/delivery"
	"github.com/quickpowered/frilly/internal/repositories/bot"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/repositories/db"
	"github.com/quickpowered/frilly/internal/repositories/serverlocations"
	"github.com/quickpowered/frilly/internal/repositories/subscriptions"
	"github.com/quickpowered/frilly/internal/repositories/web"
	"github.com/quickpowered/frilly/internal/repositories/xray"
	"github.com/quickpowered/frilly/internal/services/account"
	"github.com/quickpowered/frilly/internal/services/payment"
	"github.com/quickpowered/frilly/internal/services/promo"
	"github.com/quickpowered/frilly/internal/use-cases/commands"
	"github.com/quickpowered/frilly/internal/use-cases/commands/deps"
	expireSubscriptions "github.com/quickpowered/frilly/internal/use-cases/expire_subscriptions"
	"github.com/quickpowered/frilly/internal/use-cases/invoices"
	promoUC "github.com/quickpowered/frilly/internal/use-cases/promo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(logLevel())

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	xrayClient, err := xray.New(xrayConfigParams())
	if err != nil {
		logger.Error("failed to connect to xray", zap.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	pgPool, err := db.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Error("failed to connect to database", zap.Error(err))
		return
	}

	accountDB := db.NewAccount(pgPool)
	paymentDB := db.NewPayment(pgPool)
	promoDB := db.NewPromo(pgPool, logger)

	exchangeRates := web.NewExchangeRates()
	subscriptionsRepo, err := subscriptions.New(os.Getenv("SUBSCRIPTIONS_PATH"), exchangeRates)
	if err != nil {
		logger.Error("failed to load subscriptions", zap.Error(err))
		return
	}

	serverLocationsRepo, err := serverlocations.New()
	if err != nil {
		logger.Error("failed to load serverlocations", zap.Error(err))
		return
	}

	accountService := account.NewService(accountDB)
	paymentService := payment.NewService(paymentDB)
	promoService := promo.NewService(promoDB)

	deps := deps.NewDependencies(accountService, paymentService, promoService, subscriptionsRepo, serverLocationsRepo, xrayClient, logger)
	promoUseCase := promoUC.NewUseCase(accountService, paymentService, promoService, subscriptionsRepo, xrayClient, logger)
	invoicesUseCase := invoices.NewUseCase(accountService, paymentService, promoUseCase, subscriptionsRepo, xrayClient, logger)
	expireSubscriptionsUseCase := expireSubscriptions.NewUseCase(accountService, xrayClient, logger)
	commandsUseCase := commands.NewUseCase(deps, promoUseCase, serverLocationsRepo)

	wg := &sync.WaitGroup{}

	interval := time.Duration(30) * time.Second
	logger.Debug("starting expire subscriptions...",
		zap.Duration("interval", interval))

	wg.Go(func() {
		for {
			time.Sleep(interval)
			ctx, cancel := context.WithTimeout(context.TODO(), interval)
			expireSubscriptionsUseCase.Run(ctx)
			cancel()
		}
	})

	for _, bot := range []bin.Interface{
		bot.New("telegram", os.Getenv("TG_TOKEN")),
	} {
		wg.Go(func() {
			if err := delivery.New(bot, commandsUseCase, invoicesUseCase, logger).Listen(); err != nil {
				logger.Error("failed to start bot", zap.Error(err))
				return
			}
		})
	}

	startBanner()

	wg.Wait()
}

func logLevel() (logLevel zapcore.Level) {
	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	default:
		logLevel = zap.InfoLevel
	}
	return logLevel
}

const (
	DEFAULT_XRAY_API_ADDR    = "beyondsecure-xray:32768"
	DEFAULT_XRAY_INBOUND_TAG = "VLESS-XHTTP-REALITY"
	DEFAULT_XRAY_CONFIG_PATH = "/shared/xray/config.json"
)

func xrayConfigParams() (apiAddr string, inboundTag string, configPath string) {
	apiAddr = os.Getenv("XRAY_API_ADDR")
	if apiAddr == "" {
		apiAddr = DEFAULT_XRAY_API_ADDR
	}

	inboundTag = os.Getenv("XRAY_INBOUND_TAG")
	if inboundTag == "" {
		inboundTag = DEFAULT_XRAY_INBOUND_TAG
	}

	configPath = os.Getenv("XRAY_CONFIG_PATH")
	if configPath == "" {
		configPath = DEFAULT_XRAY_CONFIG_PATH
	}

	return apiAddr, inboundTag, configPath
}

func startBanner() {
	fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                               ║")
	fmt.Println("║               🍿 BEYOND SECURE BOT LAUNCHED 🎮                ║")
	fmt.Println("║                                                               ║")
	fmt.Println("║  ╭──────────────────────────────────────────────────────────╮ ║")
	fmt.Println("║  │  Status: ✓ Online and Ready                              │ ║")
	fmt.Println("║  │  Mode: Production                                        │ ║")
	fmt.Println("║  │  Services: Account, Payment                              │ ║")
	fmt.Println("║  │  Modules: XRay, Tools                                    │ ║")
	fmt.Println("║  ╰──────────────────────────────────────────────────────────╯ ║")
	fmt.Println("║                                                               ║")
	fmt.Println("║           Listening for incoming messages...                  ║")
	fmt.Println("║                                                               ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
}
