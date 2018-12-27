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
	// type installed struct {
	// 	ClientID     string `mapstructure:"client_id"`
	// 	ClientSecret string `mapstructure:"client_secret"`
	// 	RedirectURI  string `mapstructure:"redirect_uri"`
	// 	AuthURI      string `mapstructure:"auth_uri"`
	// 	TokenURI     string `mapstructure:"token_uri"`
	// }

	// var credentials installed

	// c := viper.Get("calendar.credentials.installed")
	err = viper.UnmarshalKey("calendar.credentials.token", &token)

	return
}
