package commands

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/telegram"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types/update"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
	"go.uber.org/zap"
)

const REPOST_CMD = "repost"

type repostHandler struct {
	deps.Dependencies
}

func NewRepostHandler(deps deps.Dependencies) *repostHandler {
	return &repostHandler{deps}
}

func (c *repostHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	if p.Message.Sender() != 924536264 {
		return bot.SendMessage(p.Message.Chat(), "You are not allowed to use this command")
	}

	if len(p.Args) < 2 {
		return bot.SendMessage(p.Message.Chat(), "Usage: /repost [message_id]")
	}

	messageID, err := strconv.Atoi(p.Args[1])
	if err != nil {
		return bot.SendMessage(p.Message.Chat(), "Invalid message ID")
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	accounts, err := c.AccountService.GetAllByPlatform(ctx, consts.PlatformTelegram)
	if err != nil {
		return err
	}

	blockedCounters := 0
	for _, accountID := range accounts {
		if err := bot.(*telegram.Adapter).ForwardMessage(update.Chat{ID: accountID}, -1003309480333, messageID); err != nil {
			if strings.Contains(err.Error(), "too many requests") {
				time.Sleep(15 * time.Second)
				continue
			}

			c.Logger.Error("failed to forward message", zap.Error(err), zap.Int("account_id", accountID))
			blockedCounters++
		}
	}

	return bot.SendMessage(p.Message.Chat(), fmt.Sprintf("Succesfully reposted!\nBlocked: %d", blockedCounters))
}
