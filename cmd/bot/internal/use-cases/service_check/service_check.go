package serviceCheck

import (
	"context"
	"strings"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types/update"
	"github.com/quickpowered/thebeyond-master/configs/language"
	"github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/quickpowered/thebeyond-master/internal/services/account"
	"go.uber.org/zap"
)

type UseCase struct {
	accountService account.Interface
	logger         *zap.Logger
}

func NewUseCase(accountService account.Interface, logger *zap.Logger) *UseCase {
	return &UseCase{accountService, logger}
}

func (uc *UseCase) Run(ctx context.Context, tgBot bin.Interface) error {
	accounts, err := uc.accountService.FindUsersForServiceCheck(ctx)
	if err != nil {
		uc.logger.Error("failed to find users for service check", zap.Error(err))
		return err
	}

	for _, account := range accounts {
		if err := uc.QueryServiceCheck(ctx, tgBot, account); err != nil {
			uc.logger.Error("failed to query service check", zap.Error(err))
			continue
		}
	}

	return nil
}

func (uc *UseCase) QueryServiceCheck(ctx context.Context, tgBot bin.Interface, account domain.ServiceCheck) error {
	if err := uc.accountService.MarkServiceCheckSent(ctx, account.ID); err != nil {
		uc.logger.Error("failed to mark service check as sent",
			zap.Error(err),
			zap.Int("account_id", account.ID),
		)
		return err
	}

	accountLanguage := strings.ToLower(strings.TrimSpace(account.Language))
	if accountLanguage == "" {
		accountLanguage = "en"
	}

	msg := i18n.ServiceCheckMessages[language.Language(accountLanguage)]
	if err := tgBot.SendMessage(update.Chat{ID: account.ExternalAccountID}, msg.ServiceWorking, &types.Keyboard{
		ButtonRows: [][]types.Button{
			{{Text: msg.YesEverythingFine, Data: "support ok"}},
			{{Text: msg.NoNeedHelp, URL: "https://t.me/beyondsecurenews?direct"}},
			{{Text: msg.ThereAreProblems, URL: "https://t.me/beyondsecurenews?direct"}},
		},
	}); err != nil {
		uc.logger.Error("failed to send service check message",
			zap.Error(err),
			zap.Int("chat_id", account.ExternalAccountID),
			zap.Int("account_id", account.ID),
		)
	}

	uc.logger.Info("service check message sent successfully",
		zap.Int("account_id", account.ID),
		zap.Int("chat_id", account.ExternalAccountID),
		zap.String("language", accountLanguage),
	)

	return nil
}
