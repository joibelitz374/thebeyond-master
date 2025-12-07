package commands

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"os"

	"github.com/quickpowered/frilly/config/language"
	"github.com/quickpowered/frilly/internal/domain"
	"github.com/quickpowered/frilly/internal/i18n"
	"github.com/quickpowered/frilly/internal/repositories/bot/bin"
	"github.com/quickpowered/frilly/internal/repositories/serverlocations"
	"github.com/quickpowered/frilly/internal/types"
	"github.com/quickpowered/frilly/internal/use-cases/commands/deps"
	"github.com/quickpowered/frilly/pkg/consts"
	"github.com/quickpowered/frilly/pkg/utils"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error { return nil }

const ACCOUNT_CMD = "account"

type accountHandler struct {
	deps      deps.Dependencies
	locations serverlocations.Repository
}

func NewAccountHandler(deps deps.Dependencies, locations serverlocations.Repository) accountHandler {
	return accountHandler{deps, locations}
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
			{{Text: "🗺 " + msg.HowToUse, Data: "howtouse"}},
			{
				{Text: "🛍 " + msg.Renew, Data: "renew"},
				{Text: "🔄 " + msg.NewKey, Data: "newkey"},
			},
			{{Text: "⛔️ " + msg.Refund, Data: "refund"}},
		},
	}
}

func (h accountHandler) activeAccountInfo(bot bin.Interface, p *domain.Payload, opts []any, msg i18n.AccountLocale) error {
	location := h.locations.Get(p.Account.ServerLocation)
	vlessURI := utils.GenerateVLESSURI(
		p.Account.KeyID,
		location.IP,
		location.SNI,
		p.Account.ShortID,
		location.Name,
	)

	buf := bytes.NewBuffer(nil)
	wr := nopCloser{Writer: buf}

	options := []standard.ImageOption{
		standard.WithQRWidth(8),
		standard.WithBorderWidth(10),
		standard.WithBuiltinImageEncoder(standard.PNG_FORMAT),
		standard.WithLogoImageFilePNG(os.Getenv("PROJECT_AVATAR_PATH")),
	}

	qrc, err := qrcode.New(vlessURI)
	if err != nil {
		return fmt.Errorf("could not generate QRCode: %v", err)
	}

	w := standard.NewWithWriter(wr, options...)
	if err := qrc.Save(w); err != nil {
		return fmt.Errorf("could not save QRCode: %v", err)
	}

	reader := bytes.NewReader(buf.Bytes())
	formattedURI := vlessURI
	if bot.GetPlatform() == consts.PlatformTelegram {
		formattedURI = "<blockquote><code>" + vlessURI + "</code></blockquote>"
	}

	remaining := utils.HumanizeDuration(utils.TimeRemaining(*p.Account.SubscriptionExpiresAt))
	if remaining == "0 minutes" {
		remaining = "waiting prolongation"
	}

	opts = append(opts, types.NewAttachments().AddFile(reader))
	return bot.SendMessage(p.Message.Chat(), fmt.Sprintf("\n🛜 %s:\n%s\n🕒 %s: %s",
		msg.Key,
		formattedURI,
		msg.Remaining,
		remaining,
	), opts...)
}
