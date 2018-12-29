package events

import (
	"net/http"

	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

// Retrieves a token, saves the token, then returns the generated client.
func GetClient(config *oauth2.Config, token *oauth2.Token) *http.Client {
	return config.Client(context.Background(), token)
}

func getToken() (token *oauth2.Token, err error) {
	err = viper.UnmarshalKey("calendar.credentials.token", &token)

	return
}
