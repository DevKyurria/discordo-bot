package http

import (
	"net/http"
	"strings"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/utils/httputil"
	"github.com/diamondburned/arikawa/v3/utils/httputil/httpdriver"
)

func NewClient(token string) *api.Client {
	token = strings.TrimSpace(token)
	if token != "" && !strings.HasPrefix(strings.ToLower(token), "bot ") {
		token = "Bot " + token
	}

	stdClient := new(http.Client)
	stdClient.Transport = NewTransport()
	httpClient := httputil.NewClientWithDriver(httpdriver.WrapClient(*stdClient))
	apiClient := api.NewCustomClient(token, httpClient)
	return apiClient
}
