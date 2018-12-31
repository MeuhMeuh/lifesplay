package events

import (
	"context"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
)

// GetClient gets the events client used to manipulate events from the global config.
func GetClient() (*http.Client, error) {
	var creds credentials
	err := viper.UnmarshalKey("calendar.credentials", &creds)
	if err != nil {
		return nil, err
	}

	var token token
	err = viper.UnmarshalKey("calendar.token", &token)

	if err != nil {
		return nil, err
	}

	oauthConfig := getOauthConfig(creds)
	oauthToken, err := getOauthToken(token)

	return oauthConfig.Client(context.Background(), oauthToken), err
}

func getOauthConfig(creds credentials) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     creds.ClientID,
		ClientSecret: creds.ClientSecret,
		RedirectURL:  creds.RedirectURI,
		Scopes:       []string{calendar.CalendarReadonlyScope},
		Endpoint: oauth2.Endpoint{
			AuthURL:  creds.AuthURI,
			TokenURL: creds.TokenURI,
		},
	}
}

func getOauthToken(tok token) (*oauth2.Token, error) {
	expiry, err := time.Parse(time.RFC3339Nano, tok.Expiry)
	if err != nil {
		return nil, err
	}

	return &oauth2.Token{
		AccessToken:  tok.AccessToken,
		TokenType:    tok.TokenType,
		RefreshToken: tok.RefreshToken,
		Expiry:       expiry,
	}, nil
}

func getToken() (token *oauth2.Token, err error) {
	err = viper.UnmarshalKey("calendar.credentials.token", &token)

	return
}
