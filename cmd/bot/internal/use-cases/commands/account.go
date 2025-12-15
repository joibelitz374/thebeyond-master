package commands

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/domain"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/i18n"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/repositories/bot/bin"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/types"
	"github.com/quickpowered/thebeyond-master/cmd/bot/internal/use-cases/commands/deps"
	"github.com/quickpowered/thebeyond-master/configs/language"
	"github.com/quickpowered/thebeyond-master/pkg/utils"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error { return nil }

const ACCOUNT_CMD = "account"

type accountHandler struct {
	deps deps.Dependencies
}

func NewAccountHandler(deps deps.Dependencies) accountHandler {
	return accountHandler{deps}
}

func (h accountHandler) Execute(bot bin.Interface, p *domain.Payload) error {
	msg := i18n.AccountMessages[language.Language(p.Account.Language)]

	text := msg.YourAccount + ":\n"
	opts := []any{deps.ToForward(bot, p), h.buildKeyboard(msg)}

	if p.Account.IsActive() {
		return h.activeAccountInfo(bot, p, opts, msg)
	}

	return bot.SendMessage(p.Message.Chat(), fmt.Sprintf("%s\n🗒 %s", text, msg.SubscriptionExpired), opts...)
}

func (h accountHandler) buildKeyboard(msg i18n.AccountLocale) *types.Keyboard {
	return &types.Keyboard{
		ButtonRows: [][]types.Button{
			{{Text: "🛜 " + msg.Connect, Data: "connect"}},
			{
				{Text: "🛍 " + msg.Renew, Data: "renew"},
				{Text: "🔄 " + msg.NewKey, Data: "newkey"},
			},
			// {
			// 	{Text: "👨‍👩‍👦‍👦 " + msg.Devices, Data: "devices"},
			// 	{Text: "💟 " + msg.Region, Data: "region"},
			// },
			{{Text: "🌈 " + msg.FREEDAYS + " 🔥", Data: "promo"}},
		},
	}
}

func (h accountHandler) activeAccountInfo(bot bin.Interface, p *domain.Payload, opts []any, msg i18n.AccountLocale) error {
	buf := bytes.NewBuffer(nil)
	wr := nopCloser{Writer: buf}

	options := []standard.ImageOption{
		standard.WithQRWidth(12),
		standard.WithBorderWidth(24),
		standard.WithBuiltinImageEncoder(standard.PNG_FORMAT),
		standard.WithLogoImageFilePNG(os.Getenv("PROJECT_AVATAR_PATH")),
	}

	subscriptionURL := "https://" + os.Getenv("PROJECT_DOMAIN") + "/sub/" + p.Account.KeyID + "/smart/ru"
	qrc, err := qrcode.New(subscriptionURL)
	if err != nil {
		return fmt.Errorf("could not generate QRCode: %v", err)
	}

	w := standard.NewWithWriter(wr, options...)
	if err := qrc.Save(w); err != nil {
		return fmt.Errorf("could not save QRCode: %v", err)
	}

	opts = append(opts, types.NewAttachments().
		AddFile(bytes.NewReader(buf.Bytes())))

	remaining := utils.HumanizeDuration(utils.TimeRemaining(*p.Account.SubscriptionExpiresAt))
	if remaining == "0 minutes" {
		remaining = "waiting prolongation"
	}

	return bot.SendMessage(p.Message.Chat(), "🕒 "+msg.Remaining+": "+remaining, opts...)
}
