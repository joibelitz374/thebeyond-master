package serviceCheck

import (
	"context"

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
		uc.logger.Error("failed to mark service check sent", zap.Error(err))
		return err
	}

	msg := i18n.ServiceCheckMessages[language.Language(account.Language)]
	if err := tgBot.SendMessage(update.Chat{ID: account.ExternalAccountID}, msg.ServiceWorking, &types.Keyboard{
		ButtonRows: [][]types.Button{
			{{Text: msg.YesEverythingFine, Data: "support ok"}},
			{{Text: msg.NoNeedHelp, URL: "https://t.me/beyondsecurebot?direct"}},
			{{Text: msg.ThereAreProblems, URL: "https://t.me/beyondsecurebot?direct"}},
		},
	}); err != nil {
		uc.logger.Error("failed to send message", zap.Error(err))
		return err
	}

	return nil
}
