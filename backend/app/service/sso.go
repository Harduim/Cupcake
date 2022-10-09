package service

import (
	"fmt"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
	"github.com/valyala/fasthttp"
)

type SSOClient struct {
	App         *confidential.Client
	ClientId    string
	RedirectUrl string
	Scopes      []string
}

func NewSSO() *SSOClient {
	return &SSOClient{}
}

func (client *SSOClient) Init(authority string, clientId string, clientSecret string, redirectUrl string, scopes []string) (error, *SSOClient) {
	cred, err := confidential.NewCredFromSecret(clientSecret)
	if err != nil {
		return fmt.Errorf("could not create a cred from a secret: %w", err), nil
	}
	confidentialClientApp, err := confidential.New(clientId,
		cred,
		confidential.WithAuthority(authority))

	client.App = &confidentialClientApp
	client.ClientId = clientId
	client.RedirectUrl = redirectUrl
	client.Scopes = scopes

	return nil, client
}

func (client *SSOClient) AuthCodeURL(ctx *fasthttp.RequestCtx) (string, error) {
	url, err := client.App.AuthCodeURL(ctx,
		client.ClientId,
		client.RedirectUrl,
		client.Scopes)

	return url, err
}
