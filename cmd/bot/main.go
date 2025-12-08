package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/delivery"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	expireSubscriptions "github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/expire_subscriptions"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/invoices"
	promoUC "github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/promo"
	"github.com/quickpowered/thebeyond-master/internal/repositories/db"
	"github.com/quickpowered/thebeyond-master/internal/repositories/subscriptions"
	"github.com/quickpowered/thebeyond-master/internal/repositories/web"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager"
	"github.com/quickpowered/thebeyond-master/internal/services/account"
	"github.com/quickpowered/thebeyond-master/internal/services/payment"
	"github.com/quickpowered/thebeyond-master/internal/services/promo"
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
	paymentDB := db.NewPayment(pgPool)
	promoDB := db.NewPromo(pgPool, logger)
	exchangeRatesRepo := web.NewExchangeRates()

	subscriptionsPath := os.Getenv("SUBSCRIPTIONS_PATH")
	if subscriptionsPath != "" {
		subscriptionsPath = "/etc/thebeyond/subscriptions/subscriptions.yaml"
	}

	subscriptionsRepo, err := subscriptions.New(subscriptionsPath, exchangeRatesRepo)
	if err != nil {
		logger.Error("failed to load subscriptions", zap.Error(err))
		return
	}

	accountService := account.NewService(accountDB)
	paymentService := payment.NewService(paymentDB)
	promoService := promo.NewService(promoDB)

	clustersPath := os.Getenv("CLUSTERS_PATH")
	if clustersPath != "" {
		clustersPath = "/etc/thebeyond/xray/clusters.yaml"
	}

	xraymanagerRepo, err := xraymanager.New(clustersPath)
	if err != nil {
		logger.Error("failed to load clusters", zap.Error(err))
		return
	}

	deps := deps.NewDependencies(accountService, paymentService, promoService, subscriptionsRepo, xraymanagerRepo, logger)
	promoUseCase := promoUC.NewUseCase(accountService, paymentService, promoService, subscriptionsRepo, logger)
	invoicesUseCase := invoices.NewUseCase(accountService, paymentService, promoUseCase, subscriptionsRepo, xraymanagerRepo, logger)
	expireSubscriptionsUseCase := expireSubscriptions.NewUseCase(accountService, xraymanagerRepo, logger)
	commandsUseCase := commands.NewUseCase(deps, promoUseCase)

	interval := time.Duration(30) * time.Second
	logger.Debug("starting expire subscriptions...", zap.Duration("interval", interval))

	wg := new(sync.WaitGroup)
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
