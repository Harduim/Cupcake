package service

import (
	"fmt"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
	"github.com/valyala/fasthttp"
)

type SSOConfig struct {
	Authority    string
	ClientId     string
	ClientSecret string
	RedirectUrl  string
	Scopes       []string
}

type SSOClient struct {
	App         *confidential.Client
	ClientId    string
	RedirectUrl string
	Scopes      []string
}

func NewSSO() *SSOClient {
	return &SSOClient{}
}

func (client *SSOClient) Init(config *SSOConfig) (error, *SSOClient) {
	cred, err := confidential.NewCredFromSecret(config.ClientSecret)
	if err != nil {
		return fmt.Errorf("could not create a cred from a secret: %w", err), nil
	}
	confidentialClientApp, err := confidential.New(config.ClientId,
		cred,
		confidential.WithAuthority(config.Authority))

	client.App = &confidentialClientApp
	client.ClientId = config.ClientId
	client.RedirectUrl = config.RedirectUrl
	client.Scopes = config.Scopes

	return nil, client
}

func (client *SSOClient) AuthCodeURL(ctx *fasthttp.RequestCtx) (string, error) {
	url, err := client.App.AuthCodeURL(ctx,
		client.ClientId,
		client.RedirectUrl,
		client.Scopes)

	return url, err
}
