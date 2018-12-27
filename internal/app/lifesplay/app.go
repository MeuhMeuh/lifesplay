package lifesplay

import (
	"net/http"
	"time"

	"github.com/meuhmeuh/lifesplay/internal/pkg/events"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
)

// App defines the app wrapper.
type App struct {
	EventsClient *http.Client
}

// Initialize sets up what's needed to run the application.
func (app *App) Initialize() {
	var creds credentials

	err := viper.UnmarshalKey("calendar.credentials", &creds)
	if err != nil {
		panic(err)
	}

	oauthConfig := &oauth2.Config{
		ClientID:     creds.ClientID,
		ClientSecret: creds.ClientSecret,
		RedirectURL:  creds.RedirectURI,
		Scopes:       []string{calendar.CalendarReadonlyScope},
		Endpoint: oauth2.Endpoint{
			AuthURL:  creds.AuthURI,
			TokenURL: creds.TokenURI,
		},
	}

	var token *token
	err = viper.UnmarshalKey("calendar.token", &token)

	expiry, err := time.Parse(time.RFC3339Nano, token.Expiry)
	if err != nil {
		panic(err)
	}

	oauthToken := &oauth2.Token{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       expiry,
	}

	app.EventsClient = events.GetClient(oauthConfig, oauthToken)
}

// Start starts the main application.
func (app *App) Start() {
	for {

	}
}
