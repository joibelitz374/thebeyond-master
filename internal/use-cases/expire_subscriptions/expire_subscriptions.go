package expireSubscriptions

import (
	"context"
	"fmt"

	"github.com/quickpowered/frilly/internal/repositories/xray"
	"github.com/quickpowered/frilly/internal/services/account"
	"go.uber.org/zap"
)

type UseCase struct {
	accountService account.Interface
	xrayClient     xray.Interface
	logger         *zap.Logger
}

func NewUseCase(accountService account.Interface, xrayClient xray.Interface, logger *zap.Logger) *UseCase {
	return &UseCase{accountService, xrayClient, logger}
}

func (uc *UseCase) Run(ctx context.Context) error {
	ids, err := uc.accountService.GetExpiredUsers(ctx)
	if err != nil {
		return err
	}

	uc.logger.Debug("expired subscriptions found", zap.Int("count", len(ids)))

	for _, id := range ids {
		email := fmt.Sprintf("id%d@user", id)

		uc.logger.Debug("expired subscription",
			zap.Int("account_id", id),
			zap.String("email", email))

		if err := uc.xrayClient.RemoveUser(email); err != nil {
			uc.logger.Error("remove user", zap.Error(err), zap.Int("account_id", id))
		}

		if err := uc.accountService.CancelSubscriptions(ctx, id); err != nil {
			uc.logger.Error("cancel subscription", zap.Error(err), zap.Int("account_id", id))
		}
	}
	return nil
}
