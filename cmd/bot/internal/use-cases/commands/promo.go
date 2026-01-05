package commands

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/telegram"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	promoCfg "github.com/quickpowered/thebeyond-master/configs/promo"
	sharedDomain "github.com/quickpowered/thebeyond-master/internal/domain"
	"github.com/quickpowered/thebeyond-master/pkg/consts"
	"github.com/quickpowered/thebeyond-master/pkg/utils"
)

const PROMO_CMD = "promo"

type promoHandler struct {
	deps.Dependencies
	promoNumberEmojis []string
}

func NewPromoCmd(deps deps.Dependencies) promoHandler {
	return promoHandler{deps, []string{"💚", "🩷", "💛"}}
}

func (h promoHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.PromoInfoMessages[p.Account.Language]
	controlMsg := i18n.ControlMessages[p.Account.Language]

	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	if len(p.Args) >= 2 {
		name := strings.ToLower(p.Args[1])
		if name == "random" {
			random, err := gonanoid.Generate(consts.NANOID_PROMO_ALPHABET, consts.NANOID_PROMO_LENGTH)
			if err != nil {
				return err
			}
			name = random
		}

		if err := h.validateName(name, msg); err != nil {
			return bot.SendMessage(p.Message.Chat(), err.Error())
		}

		promo, err := h.PromoService.Get(ctx, name)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return h.createPromo(bot, p, name, msg)
			}
			return err
		}

		shareURL := "https://t.me/share/url?url="
		shareURL += url.QueryEscape("https://t.me/beyondsecurebot?start=" + promo.Name + "\nПодключайся и пользуйся VPN вместе со мной бесплатно!")

		upperName := strings.ToUpper(promo.Name)
		text := fmt.Sprintf("🗒 %s <b>%s</b>:\n\n⭐️ <b>%s</b>: %d\n💎 <b>%s</b>: %d%%",
			msg.Instruction, upperName,
			msg.Level, promo.Level,
			msg.ClientDiscount, promoCfg.LevelDiscounts[promo.Level-1])
		text += h.statisticPromo(promo, msg)
		return bot.SendMessage(p.Message.Chat(), text, types.NewKeyboard().
			NewRow(types.NewURLButton("Share", shareURL)).
			NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, PROMO_CMD)))
	}

	promos, err := h.PromoService.GetByCreator(ctx, p.Account.ID)
	if err != nil {
		return err
	}

	text := fmt.Sprintf("%s:\n\n🆘 %s\n\n🗒 <b>%s</b>: %d/3", msg.YourPromocodes, msg.WhatIsIt, msg.Created, len(promos))
	keyboard := types.NewKeyboard()
	for i, promo := range promos {
		keyboard.NewRow(types.NewCallbackButton(h.promoNumberEmojis[i]+" "+strings.ToUpper(promo.Name), PROMO_CMD+" "+promo.Name))
	}

	if len(promos) < 3 {
		keyboard.NewRow(types.NewCallbackButton("🆕 "+msg.Create, PROMO_CMD+" random"))
	}

	keyboard.NewRow(types.NewCallbackButton("◀️ "+controlMsg.Back, MENU_CMD))
	return bot.SendMessage(p.Message.Chat(), text, keyboard)
}

func (h promoHandler) validateName(name string, msg i18n.PromoInfoLocale) error {
	if len(name) < 3 {
		return errors.New(msg.NameTooShort)
	}

	if len(name) > 16 {
		return errors.New("name must be at most 16 characters long")
	}

	for _, r := range name {
		if !(r >= 'a' && r <= 'z') {
			return errors.New(msg.NameEnglishOnly)
		}
	}

	return nil
}

func (h promoHandler) createPromo(bot bin.Interface, p *domain.Payload, name string, msg i18n.PromoInfoLocale) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	promos, err := h.PromoService.GetByCreator(ctx, p.Account.ID)
	if err != nil {
		return err
	}

	if len(promos) >= 3 {
		return bot.SendMessage(p.Message.Chat(), msg.MaxPromocodesReached)
	}

	if err := h.PromoService.Create(ctx, name, p.Account.ID); err != nil {
		return err
	}

	tgBot, ok := bot.(*telegram.Adapter)
	if !ok {
		return errors.New("failed to cast bot to telegram adapter")
	}

	if err := tgBot.AnswerCallbackQuery(p.Message.CallbackQueryID(), "🎫 "+msg.PromocodeCreated); err != nil {
		return err
	}

	p.Args = []string{PROMO_CMD, name}
	return h.Execute(bot, p)
}

func (h promoHandler) statisticPromo(promo sharedDomain.Promo, msg i18n.PromoInfoLocale) string {
	text := ""
	promoRewards := promoCfg.RewardLevels[promo.Level]

	text += fmt.Sprintf("\n\n<b>%s</b>:", msg.PromocodeClients)
	for _, count := range promoRewards.ClientTargets {
		text += formatAchievement(count, promo.Clients, promoRewards.ClientRewards[count], msg)
	}

	return text
}

func formatAchievement(target, current int, reward promoCfg.Reward, msg i18n.PromoInfoLocale) string {
	status := "🏐"
	if current >= target {
		status = "🎾"
	}

	text := fmt.Sprintf("\n• %s%s — +%d %s", utils.ToKeycaps(strconv.Itoa(target)), status, reward.Days, msg.Days)
	if reward.Discount > 0 {
		text += fmt.Sprintf(" & %d%% %s", reward.Discount, msg.RenewalDiscount)
	}

	return text
}
