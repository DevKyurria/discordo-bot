package root

import (
	"fmt"
	"strings"

	discordHttp "github.com/ayn2op/discordo/internal/http"
	"github.com/ayn2op/discordo/internal/ui/login"
	"github.com/ayn2op/tview"
)

func normalizeToken(raw string) string {
	token := strings.TrimSpace(raw)
	token = strings.TrimPrefix(token, "Bot ")
	token = strings.TrimPrefix(token, "bot ")
	return strings.TrimSpace(token)
}

type validatedTokenMsg string

func validateBotToken(rawToken string) tview.Cmd {
	return func() tview.Msg {
		token := normalizeToken(rawToken)
		if token == "" {
			return login.ErrorMsg(fmt.Errorf("missing bot token"))
		}

		me, err := discordHttp.NewClient(token).Me()
		if err != nil {
			return login.ErrorMsg(fmt.Errorf("failed to validate bot token: %w", err))
		}
		if me == nil || !me.Bot {
			return login.ErrorMsg(fmt.Errorf("this token belongs to a user account; please use a bot token from the Discord Developer Portal"))
		}

		return validatedTokenMsg(token)
	}
}
