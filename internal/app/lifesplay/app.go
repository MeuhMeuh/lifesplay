package lifesplay

import (
	"flag"
	"net/http"
	"time"

	"github.com/meuhmeuh/lifesplay/internal/pkg/events"
	"github.com/pkg/errors"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"

	astilectron "github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	astilog "github.com/asticode/go-astilog"
)

// App defines the app wrapper.
type App struct {
	EventsClient *http.Client
	Debug        *bool
	Window       *astilectron.Window
}

// Initialize sets up what's needed to run the application.
func (app *App) Initialize() {
	astilog.FlagInit()
	flag.Parse()

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
	BootGUI(app)
}

// BootGUI boots the Electron app thanks to astilectron.
func BootGUI(app *App) {
	// Init

	if err := bootstrap.Run(bootstrap.Options{
		AstilectronOptions: astilectron.Options{
			AppName:            "Lifesplay",
			AppIconDarwinPath:  "resources/logo.icns",
			AppIconDefaultPath: "resources/logo.png",
		},
		Debug: *app.Debug,
		Windows: []*bootstrap.Window{&bootstrap.Window{
			Homepage: "index.html",
			Options: &astilectron.WindowOptions{
				TitleBarStyle: astilectron.TitleBarStyleHidden,
				// Fullscreen:      astilectron.PtrBool(true), /* Will be great for the Rasp later. */
				BackgroundColor: astilectron.PtrStr("#333"),
				Center:          astilectron.PtrBool(true),
				Width:           astilectron.PtrInt(800),
				Height:          astilectron.PtrInt(480),
			},

			MessageHandler: HandleMessages,
		}},
		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astilectron.PtrStr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Label: astilectron.PtrStr("About")},
				{Role: astilectron.MenuItemRoleClose},
			},
		}},
		OnWait: func(_ *astilectron.Astilectron, iw []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			if *app.Debug {
				iw[0].OpenDevTools()
			}

			app.Window = iw[0]

			return nil
		},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}
