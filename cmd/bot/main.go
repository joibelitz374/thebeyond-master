package main

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/delivery"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/invoices"
	manageSubscriptions "github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/manage_subscriptions"
	promoUC "github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/promo"
	serviceCheck "github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/service_check"
	"github.com/quickpowered/thebeyond-master/internal/repositories/db"
	"github.com/quickpowered/thebeyond-master/internal/repositories/tariffs"
	"github.com/quickpowered/thebeyond-master/internal/repositories/web"
	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager"
	"github.com/quickpowered/thebeyond-master/internal/services/account"
	"github.com/quickpowered/thebeyond-master/internal/services/payment"
	"github.com/quickpowered/thebeyond-master/internal/services/promo"
	"github.com/quickpowered/thebeyond-master/internal/services/subscription"
	"github.com/quickpowered/thebeyond-master/pkg/banner"
	"github.com/quickpowered/thebeyond-master/pkg/logger"
	"github.com/quickpowered/thebeyond-master/pkg/postgres"
	"github.com/quickpowered/thebeyond-master/pkg/utils"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()

	logger := logger.New(os.Getenv("ENV"), os.Getenv("LOG_LEVEL"))

	pgPool, err := postgres.New(ctx, postgres.DSN())
	if err != nil {
		logger.Error("failed to connect to database", zap.Error(err))
		return
	}

	accountDB := db.NewAccount(pgPool)
	subscriptionDB := db.NewSubscription(pgPool)
	paymentDB := db.NewPayment(pgPool)
	promoDB := db.NewPromo(pgPool, logger)
	exchangeRatesRepo := web.NewExchangeRates()

	tariffsPath := os.Getenv("SUBSCRIPTIONS_PATH")
	if tariffsPath != "" {
		tariffsPath = "/etc/thebeyond/tariffs/tariffs.yaml"
	}

	tariffsRepo, err := tariffs.New(tariffsPath, exchangeRatesRepo)
	if err != nil {
		logger.Error("failed to load subscriptions", zap.Error(err))
		return
	}

	accountService := account.NewService(accountDB)
	subscriptionService := subscription.NewService(subscriptionDB)
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

	deps := deps.NewDependencies(accountService, subscriptionService, paymentService, promoService, tariffsRepo, xraymanagerRepo, logger)
	promoUseCase := promoUC.NewUseCase(accountService, subscriptionService, promoService, logger)
	invoicesUseCase := invoices.NewUseCase(accountService, subscriptionService, paymentService, promoUseCase, tariffsRepo, xraymanagerRepo, logger)
	commandsUseCase := commands.NewUseCase(deps, promoUseCase)

	manageSubscriptionsUseCase := manageSubscriptions.NewUseCase(accountService, subscriptionService, xraymanagerRepo, logger)
	serviceCheckUseCase := serviceCheck.NewUseCase(accountService, logger)

	tgBot := bot.New("telegram", os.Getenv("TG_TOKEN"))
	interval := time.Duration(30) * time.Minute
	logger.Debug("starting the automation process...")

	utils.RunPeriodic(context.Background(), interval, "reset traffic", logger, func(ctx context.Context) error {
		return manageSubscriptionsUseCase.ResetTraffic(ctx, tgBot)
	}, true)

	utils.RunPeriodic(context.Background(), interval, "disable unsub", logger, func(ctx context.Context) error {
		return manageSubscriptionsUseCase.DisableUnsub(ctx, tgBot)
	}, true)

	utils.RunPeriodic(context.Background(), interval, "disable accounts", logger, func(ctx context.Context) error {
		return manageSubscriptionsUseCase.DisableAccounts(ctx, tgBot)
	}, true)

	utils.RunPeriodic(context.Background(), time.Duration(10)*time.Minute, "enable accounts", logger, func(ctx context.Context) error {
		return manageSubscriptionsUseCase.EnableAccounts(ctx, tgBot)
	}, true)

	utils.RunPeriodic(context.Background(), time.Duration(2)*time.Minute, "service check", logger, func(ctx context.Context) error {
		return serviceCheckUseCase.Run(ctx, tgBot)
	}, true)

	wg := new(sync.WaitGroup)
	for _, bot := range []bin.Interface{tgBot} {
		wg.Go(func() {
			if err := delivery.New(bot, commandsUseCase, invoicesUseCase, logger).Listen(); err != nil {
				logger.Error("failed to start bot", zap.Error(err))
				return
			}
		})
	}

	banner.Display(os.Getenv("ENV"))

	wg.Wait()
}
