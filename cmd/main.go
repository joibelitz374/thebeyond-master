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
	"github.com/quickpowered/frilly/internal/repositories/web"
	"github.com/quickpowered/frilly/internal/repositories/xray"
	"github.com/quickpowered/frilly/internal/services/account"
	"github.com/quickpowered/frilly/internal/services/payment"
	"github.com/quickpowered/frilly/internal/use-cases/commands"
	"github.com/quickpowered/frilly/internal/use-cases/commands/tools"
	"github.com/quickpowered/frilly/internal/use-cases/invoices"
	"go.uber.org/zap"
)

const (
	DEFAULT_XRAY_API_ADDR    = "beyondsecure-warp-proxy:32768"
	DEFAULT_XRAY_INBOUND_TAG = "VLESS-XHTTP-REALITY"
	DEFAULT_XRAY_CONFIG_PATH = "/shared/xray/config.json"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	pgPool, err := db.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Error("failed to connect to database", zap.Error(err))
		return
	}

	paymentDB := db.NewPayment(pgPool)
	accountDB := db.NewAccount(pgPool)

	paymentService := payment.NewService(paymentDB)
	accountService := account.NewService(accountDB)

	apiAddr := os.Getenv("XRAY_API_ADDR")
	if apiAddr == "" {
		apiAddr = DEFAULT_XRAY_API_ADDR
	}

	inboundTag := os.Getenv("XRAY_INBOUND_TAG")
	if inboundTag == "" {
		inboundTag = DEFAULT_XRAY_INBOUND_TAG
	}

	configPath := os.Getenv("XRAY_CONFIG_PATH")
	if configPath == "" {
		configPath = DEFAULT_XRAY_CONFIG_PATH
	}

	xrayClient, err := xray.New(apiAddr, inboundTag, configPath)
	if err != nil {
		logger.Error("failed to connect to xray", zap.Error(err))
		return
	}

	modules := tools.NewModules(paymentService, accountService, xrayClient, logger)
	exchangeRates := web.NewExchangeRates()
	commandsUseCase := commands.NewUseCase(modules, exchangeRates)
	invoicesUseCase := invoices.NewUseCase(paymentService, accountService, xrayClient)

	wg := &sync.WaitGroup{}
	wg.Go(func() {
		expiredAccounts, err := accountService.GetExpiredUsers(ctx)
		if err != nil {
			logger.Error("failed to get expired users", zap.Error(err))
			return
		}

		for _, account := range expiredAccounts {
			if err := xrayClient.RemoveUser(fmt.Sprintf("id%d@user", account.ID)); err != nil {
				logger.Error("failed to remove user", zap.Error(err))
			}

			if err := accountService.SetSubscriptionExpiresAt(ctx, account.ID, nil); err != nil {
				logger.Error("failed to set subscription expires at", zap.Error(err))
			}
		}
	})

	for _, bot := range []bin.Interface{
		bot.New("telegram", os.Getenv("TG_TOKEN")),
		// bot.New("discord", os.Getenv("DS_TOKEN")),
		// bot.New("vk", os.Getenv("VK_TOKEN")),
	} {
		wg.Go(func() {
			if err := delivery.New(bot, commandsUseCase, invoicesUseCase, logger).Listen(); err != nil {
				logger.Error("failed to start bot", zap.Error(err))
				return
			}
		})
	}

	fmt.Println("\n╔═══════════════════════════════════════════════════════════════╗")
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

	wg.Wait()
}
